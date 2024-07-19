package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "{{ stabe.application_name }}",
	Short: "{{ stabe.application_description }}",
}

func init() {
	rootCmd.AddCommand(buildVersionCmd())
	// Add more commands here
}

func Execute() error {
	return rootCmd.Execute()
}
