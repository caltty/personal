package client

import (
	"testing"
)

const URL = "ldap://192.168.179.179:389"
const USERNAME = "test"
const PASSWORD = "P@ssw0rd"

func TestConn_Bind_testaccount(t *testing.T) {
	Conn_Bind(URL, USERNAME, PASSWORD)
}


func TestConn_Search(t *testing.T) {
	baseDn := "dc=hpdm,dc=com"
	Conn_Search(URL, baseDn, USERNAME, PASSWORD)
}