package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/kingpin/v2"
	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/gate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/customresourcesgate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/statemetrics"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/controller/conversion"
	authv1 "k8s.io/api/authorization/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	apisCluster "github.com/crossplane-contrib/provider-openstack/apis/cluster"
	apisNamespaced "github.com/crossplane-contrib/provider-openstack/apis/namespaced"
	"github.com/crossplane-contrib/provider-openstack/config"
	resolverapis "github.com/crossplane-contrib/provider-openstack/internal/apis"
	"github.com/crossplane-contrib/provider-openstack/internal/clients"
	controllerCluster "github.com/crossplane-contrib/provider-openstack/internal/controller/cluster"
	controllerNamespaced "github.com/crossplane-contrib/provider-openstack/internal/controller/namespaced"
	"github.com/crossplane-contrib/provider-openstack/internal/features"
)

const (
	webhookTLSCertDirEnvVar = "WEBHOOK_TLS_CERT_DIR"
	tlsServerCertDirEnvVar  = "TLS_SERVER_CERTS_DIR"
	certsDirEnvVar          = "CERTS_DIR"
	tlsServerCertDir        = "/tls/server"
)

func main() {
	var (
		app               = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for Azuread").DefaultEnvars()
		deprecationAction = func(flagName string) kingpin.Action {
			return func(c *kingpin.ParseContext) error {
				_, err := fmt.Fprintf(os.Stderr, "warning: Command-line flag %q is deprecated and no longer used. It will be removed in a future release. Please remove it from all of your configurations (ControllerConfigs, etc.).\n", flagName)
				kingpin.FatalIfError(err, "Failed to print the deprecation notice.")
				return nil
			}
		}
		debug                   = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod              = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		pollInterval            = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		pollStateMetricInterval = app.Flag("poll-state-metric", "State metric recording interval").Default("5s").Duration()
		leaderElection          = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate        = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may checked for drift from the desired state.").Default("10").Int()
		// now deprecated command-line arguments with the Terraform SDK-based upjet architecture
		_ = app.Flag("terraform-version", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Terraform version.").Envar("TERRAFORM_VERSION").Hidden().Action(deprecationAction("terraform-version")).String()
		_ = app.Flag("terraform-provider-version", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Terraform provider version.").Envar("TERRAFORM_PROVIDER_VERSION").Hidden().Action(deprecationAction("terraform-provider-version")).String()
		_ = app.Flag("terraform-native-provider-path", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Terraform native provider path for shared execution.").Default("").Envar("TERRAFORM_NATIVE_PROVIDER_PATH").Hidden().Action(deprecationAction("terraform-native-provider-path")).String()
		_ = app.Flag("terraform-provider-source", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Terraform provider source.").Envar("TERRAFORM_PROVIDER_SOURCE").Hidden().Action(deprecationAction("terraform-provider-source")).String()
		_ = app.Flag("provider-ttl", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] TTL for the native plugin processes before they are replaced. Changing the default may increase memory consumption.").Default("100").Hidden().Action(deprecationAction("provider-ttl")).Int()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()

		certsDirSet = false
		// we record whether the command-line option "--certs-dir" was supplied
		// in the registered PreAction for the flag.
		certsDir = app.Flag("certs-dir", "The directory that contains the server key and certificate.").Default(tlsServerCertDir).Envar(certsDirEnvVar).PreAction(func(_ *kingpin.ParseContext) error {
			certsDirSet = true
			return nil
		}).String()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))
	log.Default().SetOutput(io.Discard)
	ctrl.SetLogger(zap.New(zap.WriteTo(io.Discard)))

	zl := zap.New(zap.UseDevMode(*debug))
	logr := logging.NewLogrLogger(zl.WithName("provider-openstack"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	// currently, we configure the jitter to be the 5% of the poll interval
	pollJitter := time.Duration(float64(*pollInterval) * 0.05)
	logr.Debug("Starting", "sync-period", syncPeriod.String(),
		"poll-interval", pollInterval.String(), "poll-jitter", pollJitter, "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	// Get the TLS certs directory from the environment variables set by
	// Crossplane if they're available.
	// In older XP versions we used WEBHOOK_TLS_CERT_DIR, in newer versions
	// we use TLS_SERVER_CERTS_DIR. If an explicit certs dir is not supplied
	// via the command-line options, then these environment variables are used
	// instead.
	if !certsDirSet {
		// backwards-compatibility concerns
		xpCertsDir := os.Getenv(certsDirEnvVar)
		if xpCertsDir == "" {
			xpCertsDir = os.Getenv(tlsServerCertDirEnvVar)
		}
		if xpCertsDir == "" {
			xpCertsDir = os.Getenv(webhookTLSCertDirEnvVar)
		}
		// we probably don't need this condition but just to be on the
		// safe side, if we are missing any kingpin machinery details...
		if xpCertsDir != "" {
			*certsDir = xpCertsDir
		}
	}

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-openstack",
		Cache: cache.Options{
			SyncPeriod: syncPeriod,
		},
		WebhookServer: webhook.NewServer(
			webhook.Options{
				CertDir: *certsDir,
			}),
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")
	kingpin.FatalIfError(apisCluster.AddToScheme(mgr.GetScheme()), "Cannot add cluster-scoped OpenStack APIs to scheme")
	kingpin.FatalIfError(apisNamespaced.AddToScheme(mgr.GetScheme()), "Cannot add namespace-scoped OpenStack APIs to scheme")
	kingpin.FatalIfError(apiextensionsv1.AddToScheme(mgr.GetScheme()), "Cannot add api-extensions APIs to scheme")
	kingpin.FatalIfError(authv1.AddToScheme(mgr.GetScheme()), "Cannot add k8s authorization APIs to scheme")

	kingpin.FatalIfError(resolverapis.BuildScheme(apisCluster.AddToSchemes), "Cannot register the cluster-scoped OpenStack APIs with the API resolver's runtime scheme")
	kingpin.FatalIfError(resolverapis.BuildScheme(apisNamespaced.AddToSchemes), "Cannot register the namespace-scoped OpenStack APIs with the API resolver's runtime scheme")

	metricRecorder := managed.NewMRMetricRecorder()
	stateMetrics := statemetrics.NewMRStateMetrics()

	metrics.Registry.MustRegister(metricRecorder)
	metrics.Registry.MustRegister(stateMetrics)

	ctx := context.Background()

	providerCluster, err := config.GetProvider(ctx, false)
	kingpin.FatalIfError(err, "Cannot initialize the cluster-scoped provider configuration")
	providerNamespaced, err := config.GetProviderNamespaced(ctx, false)
	kingpin.FatalIfError(err, "Cannot initialize the namespace-scoped provider configuration")

	optionsCluster := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  logr,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                &feature.Flags{},
			MetricOptions: &xpcontroller.MetricOptions{
				PollStateMetricInterval: *pollStateMetricInterval,
				MRMetrics:               metricRecorder,
				MRStateMetrics:          stateMetrics,
			},
		},
		Provider:              providerCluster,
		SetupFn:               clients.TerraformSetupBuilder(providerCluster.TerraformProvider),
		PollJitter:            pollJitter,
		OperationTrackerStore: tjcontroller.NewOperationStore(logr),
		StartWebhooks:         *certsDir != "",
	}

	optionsNamespaced := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  logr,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                &feature.Flags{},
			MetricOptions: &xpcontroller.MetricOptions{
				PollStateMetricInterval: *pollStateMetricInterval,
				MRMetrics:               metricRecorder,
				MRStateMetrics:          stateMetrics,
			},
		},
		Provider:              providerNamespaced,
		SetupFn:               clients.TerraformSetupBuilder(providerNamespaced.TerraformProvider),
		PollJitter:            pollJitter,
		OperationTrackerStore: tjcontroller.NewOperationStore(logr),
		StartWebhooks:         *certsDir != "",
	}

	// https://docs.crossplane.io/latest/managed-resources/managed-resources/#managementpolicies
	if *enableManagementPolicies {
		optionsCluster.Features.Enable(features.EnableBetaManagementPolicies)
		optionsNamespaced.Features.Enable(features.EnableBetaManagementPolicies)
		logr.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	canSafeStart, err := canWatchCRD(context.TODO(), mgr)
	kingpin.FatalIfError(err, "SafeStart precheck failed")
	kingpin.FatalIfError(conversion.RegisterConversions(optionsCluster.Provider, optionsNamespaced.Provider, mgr.GetScheme()), "Cannot initialize the webhook conversion registry")
	if canSafeStart {
		crdGate := new(gate.Gate[schema.GroupVersionKind])
		optionsCluster.Gate = crdGate
		optionsNamespaced.Gate = crdGate
		kingpin.FatalIfError(customresourcesgate.Setup(mgr, xpcontroller.Options{
			Logger:                  logr,
			Gate:                    crdGate,
			MaxConcurrentReconciles: 1,
		}), "Cannot setup CRD gate")
		kingpin.FatalIfError(controllerCluster.SetupGated(mgr, optionsCluster), "Cannot setup cluster-scoped Template controllers")
		kingpin.FatalIfError(controllerNamespaced.SetupGated(mgr, optionsNamespaced), "Cannot setup namespace-scoped Template controllers")
	} else {
		logr.Info("Provider has missing RBAC permissions for watching CRDs, controller SafeStart capability will be disabled")
		kingpin.FatalIfError(controllerCluster.Setup(mgr, optionsCluster), "Cannot setup cluster-scoped OpenStack controllers")
		kingpin.FatalIfError(controllerNamespaced.Setup(mgr, optionsNamespaced), "Cannot setup namespace-scoped OpenStack controllers")
	}
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}

func canWatchCRD(ctx context.Context, mgr manager.Manager) (bool, error) {
	if err := authv1.AddToScheme(mgr.GetScheme()); err != nil {
		return false, err
	}
	verbs := []string{"get", "list", "watch"}
	for _, verb := range verbs {
		sar := &authv1.SelfSubjectAccessReview{
			Spec: authv1.SelfSubjectAccessReviewSpec{
				ResourceAttributes: &authv1.ResourceAttributes{
					Group:    "apiextensions.k8s.io",
					Resource: "customresourcedefinitions",
					Verb:     verb,
				},
			},
		}
		if err := mgr.GetClient().Create(ctx, sar); err != nil {
			return false, errors.Wrapf(err, "unable to perform RBAC check for verb %s on CustomResourceDefinitions", verbs)
		}
		if !sar.Status.Allowed {
			return false, nil
		}
	}
	return true, nil
}
