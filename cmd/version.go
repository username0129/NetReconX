package cmd

import "github.com/spf13/cobra"

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

}
