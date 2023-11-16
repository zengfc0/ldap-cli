/*
Copyright © 2023 fczeng <fczeng@fczeng.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string = "1.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get ldap-cli version",
	Long:  `get ldap-cli version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
