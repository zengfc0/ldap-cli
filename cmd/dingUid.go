/*
Copyright © 2023 fczeng <fczeng@fczeng.com>

*/
package cmd

import (
	"fmt"
	"ldap-cli/common"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// dingUidCmd represents the dingUid command
var dingUidCmd = &cobra.Command{
	Use:   "dingUid",
	Short: "get/update user dingUid from ldap",
	Long:  `get/update user dingUid from ldap.`,
	RunE:  dingUid,
}

func dingUid(cmd *cobra.Command, args []string) error {
	if os.Args[1] == "get" {
		err := getdingUid(args)
		if err != nil {
			return err
		}
	} else {
		err := updatedingUid(args)
		if err != nil {
			return err
		}
	}
	return nil
}

func getdingUid(args []string) error {
	users, err := common.GetAllUsers()

	if err != nil {
		return err
	}
	for _, user := range users {
		if strings.TrimSpace(user.Uid) == gitlabUser {
			fmt.Printf("successfully! %s's dingUid: %s\n", gitlabUser, user.DingUid)
			return nil
		}
	}
	return fmt.Errorf("cant't find user with UserUid: %s", gitlabUser)
}

func updatedingUid(args []string) error {
	if len(dingUserid) == 0 {
		return fmt.Errorf("please input the dingding userid of %s, using flag: --dinguid", gitlabUser)
	}
	err := common.UpdateAttributeValue("dingUid", gitlabUser, dingUserid)
	if err != nil {
		return err
	}
	fmt.Printf("🚀successfully update dingUid of %s\n", gitlabUser)
	return nil
}

func init() {
	getCmd.AddCommand(dingUidCmd)
	updateCmd.AddCommand(dingUidCmd)
	dingUidCmd.PersistentFlags().StringVarP(&gitlabUser, "user", "U", "", "gitlab User")
	dingUidCmd.PersistentFlags().StringVarP(&dingUserid, "dinguid", "", "", "dingding user id")
	dingUidCmd.PersistentFlags().StringVarP(&newPhoneNumber, "mobile", "M", "", "new mobile phone number")
	dingUidCmd.MarkPersistentFlagRequired("user")
}
