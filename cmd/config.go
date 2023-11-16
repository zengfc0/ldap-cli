/*
Copyright © 2023 fczeng <fczeng@fczeng.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "get user telphoneNumber from ldap",
	Long:  `get user telphoneNumber from ldap`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
