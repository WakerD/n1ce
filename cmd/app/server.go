package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func Run(runOptions *options.ServerRunOptions, stopCh <-chan struct{}) error {
	// To help debugging, immediately log version
	glog.Infof("Version: %+v", version.Get())

	server, err := CreateServerChain(runOptions, stopCh)
	if err != nil {
		return err

	}
	return server.PrepareRun().Run(stopCh)

}
