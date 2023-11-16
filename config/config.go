package config

// 全局配置变量
var Conf = new(LdapConfig)

type LdapConfig struct {
	LdapUrl    string `yaml:"ldapUrl"`
	BaseDN     string `yaml:"baseDN"`
	BindDN     string `yaml:"bindDN"`
	BindPasswd string `yaml:"bindPasswd"`
}
