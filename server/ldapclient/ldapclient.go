package ldapclient

import (
	"crypto/tls"
	"fmt"
	"log"

	"gopkg.in/ldap.v3"
)

// Bind - bind ldap server
func Bind(url string, bindusername string, bindpassword string) (*ldap.Conn, error) {
	conn, err := ldap.DialURL(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect success!")

	err = conn.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("bind success!")

	return conn, err
}

// // Bind - ldap bind
// func Bind(url string, username string, password string) {
// 	l, _ := bind(url, username, password)
// 	defer l.Close()
// }

// Search - This example demonstrates how to use the search interface
func Search(conn *ldap.Conn, baseDn string, filter string) (*ldap.SearchResult, error){

	searchRequest := ldap.NewSearchRequest(
		// "dc=example,dc=com", // The base dn to search
		baseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// "(&(objectClass=organizationalPerson))", // The filter to apply
		filter,
		[]string{"dn", "cn"}, // A list attributes to retrieve
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	return sr, err
}

// AuthByUID - auth domain account
func AuthByUID(url string, basedn string, bindusername string, bindpassword string, dn string, password string) {

	l, err := Bind(url, bindusername, bindpassword)
	defer l.Close()

	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		basedn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", dn),
		//fmt.Sprintf("(&(distinguishedName=%s))", dn), // TODO: cn or uid or tbd...
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	// Bind as the user to verify their password
	err = l.Bind(userdn, password)
	if err != nil {
		log.Fatal(err)
	}

	// Rebind as the read only user for any further queries
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
	}
}

// ModifyAttr - modify entry attribute
func ModifyAttr(url string, bindUsername string, bindPassword string, dn string, attrType string, attrVals []string) {
	l, err := Bind(url, bindUsername, bindPassword)
	defer l.Close()

	// Add a description, and replace the mail attributes
	modify := ldap.NewModifyRequest(dn, nil)
	// modify.Add("description", []string{"An test user yyyyy"})
	modify.Replace(attrType, attrVals)
	// modify.Replace("mail", []string{"user@example.org"})

	err = l.Modify(modify)
	if err != nil {
		log.Fatal(err)
	}
}

// StartTLS - This example demonstrates how to start a TLS connection
func StartTLS(url string) {
	l, err := ldap.DialURL(url)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println(l.TLSConnectionState())
	// Reconnect with TLS
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	log.Println(l.TLSConnectionState())
	if err != nil {
		log.Fatal(err)
	}

	// Operations via l are now encrypted
}
