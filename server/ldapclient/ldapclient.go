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
		log.Println(err)
		return nil, err
	}
	log.Printf("Dial url %s success!", url)

	err = conn.Bind(username, password)
	if err != nil {
		log.Println(err)
		return nil, err
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
		attrs,  // []string{"dn", "cn"}, // A list attributes to retrieve
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return sr, err
}

// AuthByUID - auth domain account
func AuthByUID(conn *ldap.Conn, baseDn string, uid string, password string) (bool, error) {

	// Search for the given username
	filter := fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", uid)
	attrs := []string{"dn"}
	sr, err := Search(conn, baseDn, filter, attrs)

	if err != nil {
		log.Println(err)
		return false, err
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	// Bind as the user to verify their password
	err = conn.Bind(userdn, password)
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Printf("Auth uid: %s (dn: %s) success!", uid, userdn)
	return true, nil
}

// AuthByDN - auth domain account
func AuthByDN(conn *ldap.Conn, baseDn string, dn string, password string) (bool, error) {
	err := conn.Bind(dn, password)
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Printf("Auth (dn: %s) success!", dn)
	return true, nil
}

// ModifyAttr - modify entry attribute
func ModifyAttr(conn *ldap.Conn, dn string, attrType string, attrVals []string) (bool, error) {

	// Add a description, and replace the mail attributes
	modifyRequest := ldap.NewModifyRequest(dn, nil)
	// modify.Add("description", []string{"An test user yyyyy"})
	modifyRequest.Replace(attrType, attrVals)
	// modify.Replace("mail", []string{"user@example.org"})

	err := conn.Modify(modifyRequest)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}

// StartTLS - This example demonstrates how to start a TLS connection
func StartTLS(url string) (*ldap.Conn, error) {
	l, err := ldap.DialURL(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer l.Close()

	log.Println(l.TLSConnectionState())
	// Reconnect with TLS
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	log.Println(l.TLSConnectionState())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return l, nil
	// Operations via l are now encrypted
}

// SearchPaging - paging
func SearchPaging(conn *ldap.Conn, baseDn string, filter string, attrs []string, pageSize uint32) {

	pagingControl := ldap.NewControlPaging(pageSize)
	controls := []ldap.Control{pagingControl}

	searchRequest := ldap.NewSearchRequest(
		baseDn, // "dc=example,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, // "(&(objectClass=organizationalPerson))", // The filter to apply
		attrs,  // []string{"dn", "cn"}, // A list attributes to retrieve
		controls,
	)

	for {

		response, err := conn.Search(searchRequest)
		if err != nil {
			log.Fatalf("Failed to execute search request: %s", err.Error())
		}

		// [do something with the response entries]
		for _, entry := range response.Entries {
			fmt.Printf("dn: %s, cn: %v\n", entry.DN, entry.GetAttributeValue("cn"))
		}
		// In order to prepare the next request, we check if the response
		// contains another ControlPaging object and a not-empty cookie and
		// copy that cookie into our pagingControl object:
		updatedControl := ldap.FindControl(response.Controls, ldap.ControlTypePaging)
		if ctrl, ok := updatedControl.(*ldap.ControlPaging); ctrl != nil && ok && len(ctrl.Cookie) != 0 {
			pagingControl.SetCookie(ctrl.Cookie)
			continue
		}
		// If no new paging information is available or the cookie is empty, we
		// are done with the pagination.
		break
	}
}
