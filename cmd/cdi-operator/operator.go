package main

import (
	"github.com/golang/glog"

	"kubevirt.io/containerized-data-importer/pkg/operator"
)

func printVersion() {
	glog.V(1).Infof(fmt.Sprintf("Go Version: %s", runtime.Version()))
	glog.V(1).Infof(fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH))
	glog.V(1).Infof(fmt.Sprintf("operator-sdk Version: %v", sdkVersion.Version))
}

func main() {

	printVersion()
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		glog.V(1).Fatalf(err, "failed to get watch namespace")
	}

	// glog.V(2).Infoln("cdi controller exited")

	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{Namespace: namespace})
	if err != nil {
		glog.V(1).Fatalf(err, "")
	}

	// log.Info("Registering Components.")

	// Setup Scheme for all resources
	if err := operator.AddToScheme(mgr.GetScheme()); err != nil {
		glog.V(1).Fatalf(err, "")
	}

	// Setup all Controllers
	if err := operator.AddToManager(mgr); err != nil {
		glog.V(1).Fatalf(err, "")
	}

	// log.Info("Starting the Cmd.")

	// Start the Cmd
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		glog.V(1).Fatalf(err, "manager exited non-zero")
	}
}
