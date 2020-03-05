# Go-LDAP demo
This is a sample demo to show how integrate ldap api within go to connect a AD or other LDAP server.

## Dependencies
- [go-ldap/ldap](https://github.com/go-ldap/ldap/tree/master/v3)

## TODO
### Requriement level
- API to be implemented
  - [x] Bind
  - [x] Search
  - [ ] Authenticate
  
- [ ] SSL support

### Coding Level
- ~~how to reuse method (for the defer concern)~~
  - each defer will be invoked after method. Fix - need to provide private api


## Other Tools
- Apache Directory
    - [How to add a new user?](http://opendesignarch.blogspot.com/2012/12/adding-new-user-to-apacheds-using.html)