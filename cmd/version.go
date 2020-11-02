package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Long:  "Prints the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v: %v\n", color.MagentaString("Stylus Version"), color.GreenString(VERSION))
	},
}
