package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "Server is the backend of the NetReconX project",
		Run: func(cmd *cobra.Command, args []string) {
			tip()
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)
}

func tip() {

}

func Execute() error {
	return rootCmd.Execute()
}
