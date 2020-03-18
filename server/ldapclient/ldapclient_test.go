package ldapclient

import (
	"testing"
)

// Windows AD Server
const (
	url          = "ldap://192.168.179.182:389"
	bindUsername = "administrator@hpdm.com"
	bindPassword = "P@ssw0rd"
	baseDn       = "dc=hpdm,dc=com"

	authUserDn   = "CN=test,CN=Users,DC=hpdm,DC=com"
	authPassword = "P@ssw0rd"

	userDn = "CN=test,CN=Users,DC=hpdm,DC=com"
)

// Apache Directory LDAP Server
// const (
// 	URL           = "ldap://localhost:10389"
// 	BIND_USERNAME = "uid=admin,ou=system"
// 	BIND_PASSWORD = "secret"
// 	BASE_DN       = "dc=example,dc=com"

// 	AUTH_USER_DN = "cn=Jason Shi,ou=users,dc=example,dc=com"
// 	AUTH_PASSWORD = "test"

// 	USER_DN = "employeeNumber=1,ou=users,dc=example,dc=com"
// )

func Test_Bind_testaccount(t *testing.T) {
	Bind(url, bindUsername, bindPassword)
}

func Test_Search_Person(t *testing.T) {
	filterPerson := "(&(objectClass=organizationalPerson))"
	Search(url, baseDn, bindUsername, bindPassword, filterPerson)
}

func Test_Auth(t *testing.T) {
	Auth(url, baseDn, bindUsername, bindPassword, authUserDn, authPassword)
}

func Test_Modify(t *testing.T) {
	attrType := "description"
	attrVals := []string{"An test user description"}
	ModifyAttr(url, bindUsername, bindPassword, userDn, attrType, attrVals)
}

func Test_StartTLS(t *testing.T) {
	StartTLS(url)
}
