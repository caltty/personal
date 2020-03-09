package client

import (
	"testing"
)

// Windows AD Server
// const (
// 	URL           = "ldap://192.168.179.182:389"
// 	BIND_USERNAME = "administrator@hpdm.com"
// 	BIND_PASSWORD = "P@ssw0rd"
// 	BASE_DN       = "dc=hpdm,dc=com"
//  AUTH_USERNAME = "test"
//  AUTH_PASSWORD = "P@ssw0rd"
//  USER_DN = "CN=test,CN=Users,DC=hpdm,DC=com"
// )


// Apache Directory LDAP Server
const (
	URL           = "ldap://localhost:10389"
	BIND_USERNAME = "uid=admin,ou=system"
	BIND_PASSWORD = "secret"
	BASE_DN       = "dc=example,dc=com"

	AUTH_UID = "jasons"
	AUTH_PASSWORD = "test"

	USER_DN = "employeeNumber=1,ou=users,dc=example,dc=com"
)

func TestConn_Bind_testaccount(t *testing.T) {
	Conn_Bind(URL, BIND_USERNAME, BIND_PASSWORD)
}

func TestConn_Search_Person(t *testing.T) {
	filter_person := "(&(objectClass=organizationalPerson))"
	Conn_Search(URL, BASE_DN, BIND_USERNAME, BIND_PASSWORD, filter_person)
}

func Test_Auth(t *testing.T) {
	Auth(URL, BASE_DN, BIND_USERNAME, BIND_PASSWORD, AUTH_UID, AUTH_PASSWORD)
}

func Test_Conn_Modify(t *testing.T) {
	attrType := "description"
	attrVals := []string{"An test user description"}
	Conn_Modify_Attr(URL, BIND_USERNAME, BIND_PASSWORD, USER_DN, attrType, attrVals)
}

func Test_StartTLS(t *testing.T) {
	StartTLS(URL)
}
