package client

import (
	"gopkg.in/ldap.v3"
	"log"
)

func Conn_Bind() {
	l, err := ldap.DialURL("ldap://localhost:10389")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect success!")
	defer l.Close()

	err = l.Bind("uid=jason,ou=users,ou=system", "jason")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("bind success!")
}