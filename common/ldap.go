package common

import (
	"fmt"
	"ldap-cli/config"
	"net"
	"strings"
	"time"

	ldap "github.com/go-ldap/ldap/v3"
)

type User struct {
	DN              string `json:"dn"`
	Uid             string `json:"uid"`
	CN              string `json:"cn"`
	SN              string `json:"sn"`
	TelephoneNumber string `json:"telephoneNumber"`
	DisplayName     string `json:"displayName"`    // 展示名字，可以是中文名字
	Mail            string `json:"mail"`           // 邮箱
	EmployeeNumber  string `json:"employeeNumber"` // 员工工号
	DingUid         string `json:"dingUid"`
}

func InitLDAPConn() (*ldap.Conn, error) {
	ldapConn, err := ldap.DialURL(config.Conf.LdapUrl, ldap.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}))
	if err != nil {
		return nil, err
	}

	err = ldapConn.Bind(config.Conf.BindDN, config.Conf.BindPasswd)
	if err != nil {
		return nil, err
	}
	return ldapConn, err
}

func GetAllUsers() (users []*User, err error) {
	searchRequest := ldap.NewSearchRequest(
		config.Conf.BaseDN,                                          // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		"(&(objectClass=organizationalPerson))", // This is Filter for LDAP query
		[]string{},                              // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
		nil,
	)

	Conn, err := InitLDAPConn()
	defer Conn.Close()
	if err != nil {
		return nil, err
	}

	sr, err := Conn.Search(searchRequest)
	if err != nil {
		return users, err
	}

	if len(sr.Entries) > 0 {
		for _, v := range sr.Entries {
			users = append(users, &User{
				DN:              v.DN,
				Uid:             v.GetAttributeValue("uid"),
				CN:              v.GetAttributeValue("cn"),
				SN:              v.GetAttributeValue("sn"),
				TelephoneNumber: v.GetAttributeValue("telephoneNumber"),
				DisplayName:     v.GetAttributeValue("displayName"),
				Mail:            v.GetAttributeValue("mail"),
				EmployeeNumber:  v.GetAttributeValue("employeeNumber"),
				DingUid:         v.GetAttributeValue("dingUid"),
			})
		}
	}
	return
}

func GetUserDN(UserUid string) (UserDN string, err error) {
	users, err := GetAllUsers()

	if err != nil {
		return "", err
	}
	for _, user := range users {
		if strings.TrimSpace(user.Uid) == UserUid {
			return user.DN, nil
		}
	}
	return "", fmt.Errorf("cant't find user with UserUid: %s", UserUid)
}

func UpdateAttributeValue(attribute string, userId string, newAttributeValue string) error {
	userDN, err := GetUserDN(userId)

	if err != nil {
		return err
	}

	modify := ldap.NewModifyRequest(userDN, nil)
	modify.Replace(string(attribute), []string{newAttributeValue})

	conn, err := InitLDAPConn()
	defer conn.Close()
	if err != nil {
		return err
	}
	err = conn.Modify(modify)
	if err != nil {
		return err
	}
	return nil
}
