package ldapclient

import (
	"crypto/tls"
	"fmt"
	"log"

	"gopkg.in/ldap.v3"
)

// Bind - bind ldap server
func Bind(url string, username string, password string) (*ldap.Conn, error) {
	conn, err := ldap.DialURL(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Dial url %s success!", url)

	err = conn.Bind(username, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bind %s success!", username)

	return conn, err
}

// Search - search from base dn, return with specified attributes
func Search(conn *ldap.Conn, baseDn string, filter string, attrs []string) (*ldap.SearchResult, error) {

	searchRequest := ldap.NewSearchRequest(
		baseDn, // "dc=example,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, // "(&(objectClass=organizationalPerson))", // The filter to apply
		attrs, // []string{"dn", "cn"}, // A list attributes to retrieve
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	return sr, err
}

// AuthByUID - auth domain account
func AuthByUID(conn *ldap.Conn, baseDn string, uid string, password string) {

	// Search for the given username
	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", uid)
	attrs := []string{"dn"}
	sr, err := Search(conn, baseDn, filter, attrs)

	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	// Bind as the user to verify their password
	err = conn.Bind(userdn, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Auth uid: %s (dn: %s) success!", uid, userdn)
}

// ModifyAttr - modify entry attribute
func ModifyAttr(conn *ldap.Conn, dn string, attrType string, attrVals []string) {

	// Add a description, and replace the mail attributes
	modifyRequest := ldap.NewModifyRequest(dn, nil)
	// modify.Add("description", []string{"An test user yyyyy"})
	modifyRequest.Replace(attrType, attrVals)
	// modify.Replace("mail", []string{"user@example.org"})

	err := conn.Modify(modifyRequest)
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
