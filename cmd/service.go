package cmd

import (
	"github.com/spf13/cobra"
)

var ServiceName string
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Specify a service that needs to be started",
	Run: func(cmd *cobra.Command, args []string) {
		panic(ServiceName)
	},
}

func init() {
	rootCmd.AddCommand(ServiceCmd)

	ServiceCmd.Flags().StringVar(&ServiceName, "name", "", "service name")
}
