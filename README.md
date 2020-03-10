# Go-LDAP/AD demo
This is a sample demo to show how integrate ldap api within go to connect a AD or other LDAP server.

## Dependencies
- [go-ldap/ldap](https://github.com/go-ldap/ldap/tree/master/v3)
- Setting up AD/Server
  - [x] Windows AD Server 
  - [x] Apache Directory Server


## TODO
### Requriement level
- API to be implemented
  - [x] Bind
    - [x] unit test
  - [x] Search
    - [x] unit test
  - [ ] Authenticate
    - [x] auth with cn
    - [ ] auth with dn
    - [ ] auth with s
    - [ ] unit test
  - [x] Modify attribute
    - [x] unit test
  
- [ ] SSL/TLS support
  - [ ] Windows AD Server 

  - [ ] Apache Directory Server
    - [x] Start TLS
    - [x] By default, Apache Directory will create self-signed certficate
    -   For TLS support, just go to Cofiguration/LDAPS Server/Advanced -> Enable TLS would be okay
      - No need to change port
    - [ ] TLS details



### Design & Maintainability Level
- More Object-Oriented
  - [ ] Abstract struct/interface
    - [ ] Binding struct
    - [ ] Searching struct
    - [ ] Searching interface

### Problem solved (coding level)
- ~~how to reuse method (for the defer concern)~~
  - [x] each defer will be invoked after method. Fix - need to provide private api

- ~~Rename symbols of go always fails~~
  - [x] works in my another vscode env
  - [x] [Go Please! Visual Studio Code + Go Mod + Go Language Server](http://www.matthiassommer.it/programming/go-please-visual-studio-code-go-mod-go-language-server/)
    - "go.useLanguageServer": true
    - setx GO111MODULE on

- ~~go comment has specified usage~~
  - [x] `// +build linux` for building for linux 




## Env & Tools
- Setting up LDAP/AD Server
  - AD Sever (Windows)
    - [Windows Server 2019 – Active Directory Installation Beginners Guide](https://www.moderndeployment.com/windows-server-2019-active-directory-installation-beginners-guide/)
  - [Apache Directory](https://directory.apache.org/)
    - [How to add a new user?](http://opendesignarch.blogspot.com/2012/12/adding-new-user-to-apacheds-using.html)