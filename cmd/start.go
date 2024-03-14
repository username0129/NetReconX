package cmd

import (
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the Gin web server",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func init() {
	startCmd.Flags()
}

func start() {

}
