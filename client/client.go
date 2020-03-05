package client

import (
	"fmt"
	"gopkg.in/ldap.v3"
	"log"
)

func conn_bind(url string, username string, password string) *ldap.Conn {
	l, err := ldap.DialURL(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect success!")

	err = l.Bind(username, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("bind success!")

	return l
}

func Conn_Bind(url string, username string, password string) {
	l := conn_bind(url, username, password)
	defer l.Close()
}

// This example demonstrates how to use the search interface
func Conn_Search(url string, baseDn string, username string, password string) {
	l := conn_bind(url, username, password)
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		// "dc=example,dc=com", // The base dn to search
		baseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))", // The filter to apply
		[]string{"dn", "cn"},                    // A list attributes to retrieve
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
