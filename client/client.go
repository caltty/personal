package client

import (
	"fmt"
	"gopkg.in/ldap.v3"
	"log"
)

func conn_bind(url string, bindusername string, bindpassword string) (*ldap.Conn, error) {
	l, err := ldap.DialURL(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect success!")

	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("bind success!")

	return l, err
}

func Conn_Bind(url string, username string, password string) {
	l, _ := conn_bind(url, username, password)
	defer l.Close()
}

// This example demonstrates how to use the search interface
func Conn_Search(url string, baseDn string, bindUsername string, bindPassword string, filter string) {
	l, err := conn_bind(url, bindUsername, bindPassword)
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		// "dc=example,dc=com", // The base dn to search
		baseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// "(&(objectClass=organizationalPerson))", // The filter to apply
		filter,
		[]string{"dn", "cn"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}

// auth domain account
func Auth(url string, basedn string, bindusername string, bindpassword string, username string, password string) {

	l, err := conn_bind(url, bindusername, bindpassword)
	defer l.Close()

	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		basedn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		fmt.Sprintf("(&(objectClass=organizationalPerson)(cn=%s))", username), // TODO: cn or uid or tbd...
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

func Conn_Modify(url string, bindUsername string, bindPassword string, dn string) {
	l, err := conn_bind(url, bindUsername, bindPassword)
	defer l.Close()

	// Add a description, and replace the mail attributes
	modify := ldap.NewModifyRequest(dn, nil)
	// modify.Add("description", []string{"An test user yyyyy"})
	modify.Replace("description", []string{"An test user zzzz"})
	// modify.Replace("mail", []string{"user@example.org"})

	err = l.Modify(modify)
	if err != nil {
		log.Fatal(err)
	}
}
