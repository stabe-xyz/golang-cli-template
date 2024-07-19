package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cobraCmd *cobra.Command, args []string) {
		fmt.Println("{{ stabe.application_name }} v0.0.1")
	},
}

func buildVersionCmd() *cobra.Command {
	return versionCmd
}
