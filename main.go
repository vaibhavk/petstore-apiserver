package main

import (
	"flag"

	"k8s.io/klog"

	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/logs"

	"github.com/vaibhavk/petstore-apiserver/pkg/cmd/server"

	// load all auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	stopCh := genericapiserver.SetupSignalHandler()
	options := server.NewCustomServerOptions()
	cmd := server.NewCommandStartCustomServer(options, stopCh)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
