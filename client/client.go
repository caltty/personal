package main

import (
	"gopkg.in/ldap.v3"
	"log"
)

func Conn_Bind() {
	l, err := ldap.DialURL("ldap://localhost:10389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind("uid=admin,ou=system", "secret")
	if err != nil {
		log.Fatal(err)
	}
}
