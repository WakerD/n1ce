package app

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"n1ce/server/models"
	"n1ce/server/router"
)

func init() {
	RootCmd.AddCommand(startCmd)
}

var (
	RootCmd = &cobra.Command{
		Use:   "n1ce",
		Short: "a basic http server",
		Long:  `a project for play, u can join me`,
	}
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start server on port",
		Long:  `the project will runing`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server start ...")
			models.ConnectMongodb("localhost", "test")
			models.ConnectRedis()
			if len(args) == 0 {
				http.ListenAndServe(":3737", router.Load())
			} else {
				http.ListenAndServe(":"+args[0], router.Load())
			}
		},
	}
)

func Run() error {
	err := RootCmd.Execute()
	return err
}
