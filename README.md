## PHP LDAP Admin login

### org0

- URL: https://localhost:7443/
- Login DN: dc=example
- Password: admin0

### org1

- URL: https://localhost:8443/
- Login DN: dc=example
- Password: admin0

### person

- URL: https://localhost:9443/
- Login DN: dc=example
- Password: admin0


## Services
### httpd

``` 
curl -u taro.yamada:taro "http://localhost:8081/internal/"
```

- account: taro.yamada
- password: taro

note that jiro.sato does not have an access to the internal site

### sshd

```
ssh -p 8022 jiro.sato@localhost
```

- account: jiro.sato
- password: jiro

note that taro.yamada does not have an access to sshd


### address

host name: localhost
base identifier: dc=org0,dc=example
port: 389
