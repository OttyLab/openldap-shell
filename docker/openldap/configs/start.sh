#!/bin/sh

/app/libexec/slapd -h "ldap:///openldap ldapi:///"

/app/bin/ldapadd -x -D "cn=Manager,dc=example,dc=com" -w admin0 -f /configs/init.ldif
sleep 1

/app/bin/ldapmodify -Q -Y EXTERNAL -H ldapi:/// -f /configs/add-configpw.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/cosine.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/inetorgperson.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/nis.ldif


while :
do
    sleep 10
done
