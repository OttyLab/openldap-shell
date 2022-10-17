#!/bin/sh

/app/libexec/slapd -h "ldap:///openldap ldapi:///"

/app/bin/ldapmodify -Q -Y EXTERNAL -H ldapi:/// -f /configs/add-configpw.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/cosine.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/inetorgperson.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/nis.ldif


while :
do
    sleep 10
done
