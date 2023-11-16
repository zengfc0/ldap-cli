/*
Copyright © 2023 fczeng <fczeng@fczeng.com>

*/
package cmd

import (
	"fmt"
	"ldap-cli/common"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: search,
}

func search(cmd *cobra.Command, args []string) error {
	users, err := common.GetAllUsers()
	if err != nil {
		return err
	}
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
