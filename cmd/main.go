package main

import (
	//"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	//	"github.com/spf13/pflag"

	"n1ce/cmd/app"
	//	"n1ce/cmd/app/options"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//s := options.NewServerRunOptions()
	//s.AddFlags(pflag.CommandLine)

	//flag.InitFlags()

	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
