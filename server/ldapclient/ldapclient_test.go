package ldapclient

import (
	"fmt"
	"log"
	"testing"
)

// Windows AD Server
// const (
// 	url          = "ldap://192.168.179.182:389"
// 	bindUsername = "administrator@hpdm.com"
// 	bindPassword = "P@ssw0rd"
// 	baseDn       = "dc=hpdm,dc=com"

// 	authUserDn  = "CN=test,CN=Users,DC=hpdm,DC=com"
// 	authPassword = "P@ssw0rd"

// 	userDn = "CN=test,CN=Users,DC=hpdm,DC=com"
// )

// Apache Directory LDAP Server
const (
	url          = "ldap://localhost:10389"
	bindUsername = "uid=admin,ou=system"
	bindPassword = "secret"
	baseDn       = "dc=shishuwu,dc=com"

	testDn       = "uid=test,ou=users,dc=shishuwu,dc=com"
	testPassword = "secret"
	testUID      = "test"

	userDn = "uid=test,ou=users,dc=shishuwu,dc=com"
)

func Test_Bind(t *testing.T) {
	conn, err := Bind(url, bindUsername, bindPassword)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(conn)
}

func Test_Search_Person(t *testing.T) {
	conn, _ := Bind(url, bindUsername, bindPassword)
	defer conn.Close()

	filterPerson := "(&(objectClass=organizationalPerson))"
	sr, _ := Search(conn, baseDn, filterPerson)

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}

func Test_AuthByUid(t *testing.T) {
	AuthByUID(url, baseDn, bindUsername, bindPassword, testUID, testPassword)
}

func Test_Modify(t *testing.T) {
	attrType := "description"
	attrVals := []string{"An test user description"}
	ModifyAttr(url, bindUsername, bindPassword, userDn, attrType, attrVals)
}

func Test_StartTLS(t *testing.T) {
	StartTLS(url)
}
