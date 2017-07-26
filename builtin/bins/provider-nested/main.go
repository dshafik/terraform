package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/hashicorp/terraform/builtin/providers/nested"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	pid := os.Getpid()
	f, err := os.Create(fmt.Sprintf("/tmp/tf-nested-cpu-%d.prof", pid))
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nested.Provider,
	})
}
