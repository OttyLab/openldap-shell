## PHP LDAP Admin login

- URL: https://localhost:6443/
- Login DN: dc=root
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
base identifier: ou=org1,dc=root
port: 389
