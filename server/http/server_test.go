package server

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

const (
	test2DN       = "CN=testuser2,OU=argonldap,OU=shanghai,DC=sh,DC=argon"
	test2Password = "Shanghai2010"
)

var respMap map[string]interface{} = make(map[string]interface{})
var token string

func Test_Login(t *testing.T) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"username": test2DN,
		"password": test2Password,
	})
	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)
	client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	resp, err := client.Post("https://localhost:9443/login", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
	json.Unmarshal(respBody, &respMap)
	if err != nil {
		log.Fatal(err)
	}
	token = respMap["token"].(string)
}

func Test_Modify_Attributes(t *testing.T) {
	Test_Login(t)

	requestBody, err := json.Marshal(map[string]interface{}{
		"dn":       test2DN,
		"attrType": "description",
		"attrVals": []string{"a test descriptionaa"},
	})
	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)
	client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	req, err := http.NewRequest("Post", "https://localhost:9443/modify-attribute", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}
