package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"CampusRecruitment/pkg/consts"
)

var rootCmd = &cobra.Command{
	Use:   "CampusRecruitment",
	Short: "CampusRecruitment",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s, commit: %s\n", consts.VERSION, consts.COMMIT)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serveCmd)
}
