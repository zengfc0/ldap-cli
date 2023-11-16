/*
Copyright © 2023 fczeng <fczeng@fczeng.com>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	//"github.com/fczeng0/ldap-cli/config"
	"ldap-cli/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	LdapUrl    string
	BindDN     string
	BindPasswd string

	gitlabUser     string
	newPhoneNumber string
	dingUserid     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ldap-cli",
	Short: "A client for LDAP directory servers.",
	Long:  `A client for LDAP directory servers.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ldap.yaml)")
	rootCmd.PersistentFlags().StringVarP(&LdapUrl, "ldapurl", "H", "", "LDAP Uniform Resource Identifier(s)")
	rootCmd.PersistentFlags().StringVarP(&BindDN, "binddn", "D", "", "bind DN")
	rootCmd.PersistentFlags().StringVarP(&BindPasswd, "bindpasswd", "P", "", "bind password (for simple authentication)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ldap")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	notFound := &viper.ConfigFileNotFoundError{}
	switch {
	case err != nil && !errors.As(err, notFound):
		cobra.CheckErr(err)
	case err != nil && errors.As(err, notFound):
		// The config file is optional, we shouldn't exit when the config is not found
		break
	default:
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	if LdapUrl != "" {
		viper.Set("ldapurl", LdapUrl)
	}

	if BindDN != "" {
		viper.Set("binddn", BindDN)
	}

	if BindPasswd != "" {
		viper.Set("bindpasswd", BindPasswd)
	}

	if err := viper.Unmarshal(config.Conf); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s", err))
	}
}
