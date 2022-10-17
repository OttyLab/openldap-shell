#!/bin/sh

/app/libexec/slapd -h "ldap:///openldap ldapi:///"
sleep 1
/app/bin/ldapadd -x -D "cn=Manager,dc=root" -w admin0 -f /configs/init.ldif
sleep 1
/app/bin/ldapadd -x -D "cn=Manager,dc=root" -w admin0 -f /configs/org1.ldif
sleep 1
/app/bin/ldapadd -x -D "cn=Manager,dc=root" -w admin0 -f /configs/org2.ldif
sleep 1

/app/bin/ldapmodify -Q -Y EXTERNAL -H ldapi:/// -f /configs/add-configpw.ldif
sleep 1
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/cosine.ldif
sleep 1
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/inetorgperson.ldif
sleep 1
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/nis.ldif
sleep 1

/app/bin/ldapadd -x -D "cn=Manager,dc=root" -w admin0 -f /configs/user-groups.ldif
sleep 1
/app/bin/ldapadd -x -D "cn=Manager,dc=root" -w admin0 -f /configs/users.ldif
sleep 1


while :
do
    sleep 10
done
