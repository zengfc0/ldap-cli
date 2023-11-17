/*
Copyright © 2023 fczeng <fczeng@fczeng.com>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"ldap-cli/common"

	"github.com/spf13/cobra"
)

// phoneNumberCmd represents the phoneNumber command
var phoneNumberCmd = &cobra.Command{
	Use:   "phoneNumber",
	Short: "get/update user telphoneNumber from ldap",
	Long:  `get/update user telphoneNumber from ldap`,
	RunE:  phoneNumber,
}

func phoneNumber(cmd *cobra.Command, args []string) error {
	if os.Args[1] == "get" {
		err := getPhoneNumber(args)
		if err != nil {
			return err
		}
	} else {
		err := updatePhoneNumber(args)
		if err != nil {
			return err
		}
	}
	return nil
}

func getPhoneNumber(args []string) error {
	users, err := common.GetAllUsers()

	if err != nil {
		return err
	}
	for _, user := range users {
		if strings.TrimSpace(user.Uid) == gitlabUser {
			fmt.Printf("successfully! %s's telephoneNumber: %s\n", gitlabUser, user.TelephoneNumber)
			return nil
		}
	}
	return fmt.Errorf("cant't find user with UserUid: %s", gitlabUser)
}

func updatePhoneNumber(args []string) error {
	if len(newPhoneNumber) == 0 {
		return fmt.Errorf("please input the telephone number of %s, using flag: -M", gitlabUser)
	}
	err := common.UpdateAttributeValue("telephoneNumber", gitlabUser, newPhoneNumber)
	if err != nil {
		return err
	}
	fmt.Printf("🚀successfully update phoneNumber of %s\n", gitlabUser)
	return nil
}

func init() {
	getCmd.AddCommand(phoneNumberCmd)
	updateCmd.AddCommand(phoneNumberCmd)
	phoneNumberCmd.PersistentFlags().StringVarP(&gitlabUser, "user", "U", "", "gitlab User")
	phoneNumberCmd.PersistentFlags().StringVarP(&newPhoneNumber, "mobile", "M", "", "new mobile phone number")
	phoneNumberCmd.MarkPersistentFlagRequired("user")
}
