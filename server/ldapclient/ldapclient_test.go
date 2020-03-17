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
	// url_ssl      = "ldaps://localhost:10636"
	
	bindUsername = "uid=admin,ou=system"
	bindPassword = "secret"
	baseDn       = "dc=shishuwu,dc=com"

	testDn       = "uid=test,ou=users,dc=shishuwu,dc=com"
	testPassword = "secret"
	testUID      = "test"

	jasonSAMAccountName = "jasons"

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
	attrs := []string{"dn", "cn"}
	sr, _ := Search(conn, baseDn, filterPerson, attrs)

	for _, entry := range sr.Entries {
		fmt.Printf("dn: %s, cn: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}

func Test_Search_Person_sAMAccountName(t *testing.T) {
	conn, _ := Bind(url, bindUsername, bindPassword)
	defer conn.Close()

	filterPerson := fmt.Sprintf("(&(objectClass=organizationalPerson)(sAMAccountName=%s))", jasonSAMAccountName)
	attrs := []string{"dn", "cn"}
	sr, _ := Search(conn, baseDn, filterPerson, attrs)

	for _, entry := range sr.Entries {
		fmt.Printf("dn: %s, cn: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}

func Test_Search_Person_Paging(t *testing.T) {
	conn, _ := Bind(url, bindUsername, bindPassword)
	defer conn.Close()

	filterPerson := "(&(objectClass=organizationalPerson))"
	attrs := []string{"dn", "cn"}
	SearchPaging(conn, baseDn, filterPerson, attrs, 1)

}


func Test_Search_Person_UID(t *testing.T) {
	conn, _ := Bind(url, bindUsername, bindPassword)
	defer conn.Close()

	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", testUID)
	attrs := []string{"dn", "cn"}
	sr, _ := Search(conn, baseDn, filter, attrs)

	for _, entry := range sr.Entries {
		fmt.Printf("dn: %s, cn: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}


func Test_AuthByUid(t *testing.T) {
	conn, _ := Bind(url, bindUsername, bindPassword)
	defer conn.Close()

	AuthByUID(conn, baseDn, testUID, testPassword)
}

func Test_Modify(t *testing.T) {
	conn, _ := Bind(url, bindUsername, bindPassword)
	defer conn.Close()

	attrType := "description"
	attrVals := []string{"An test user description"}
	ModifyAttr(conn, userDn, attrType, attrVals)
}

func Test_StartTLS(t *testing.T) {
	StartTLS(url)
}
