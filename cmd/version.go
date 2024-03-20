package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"server/internal/global"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of server",
		Run: func(cmd *cobra.Command, args []string) {
			version()
		},
	}
)

func version() {
	fmt.Printf("ServerConfig version %s -- HEAD\n", global.Version)
}
