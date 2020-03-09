package client

import (
	"testing"
)

const (
	URL           = "ldap://192.168.179.182:389"
	BIND_USERNAME = "administrator@hpdm.com"
	BIND_PASSWORD = "P@ssw0rd"
	BASE_DN       = "dc=hpdm,dc=com"
)

func TestConn_Bind_testaccount(t *testing.T) {
	Conn_Bind(URL, BIND_USERNAME, BIND_PASSWORD)
}

func TestConn_Search_Person(t *testing.T) {
	filter_person := "(&(objectClass=organizationalPerson))"
	Conn_Search(URL, BASE_DN, BIND_USERNAME, BIND_PASSWORD, filter_person)
}

func Test_Auth(t *testing.T) {
	username := "test"
	password := "P@ssw0rd"

	Auth(URL, BASE_DN, BIND_USERNAME, BIND_PASSWORD, username, password)
}

func Test_Conn_Modify(t *testing.T) {
	testDn := "CN=test,CN=Users,DC=hpdm,DC=com"
	attrType := "description"
	attrVals := []string{"An test user description"}
	Conn_Modify_Attr(URL, BIND_USERNAME, BIND_PASSWORD, testDn, attrType, attrVals)
}
