#!/bin/sh

/app/libexec/slapd -h "ldap:///openldap ldapi:///"
/app/bin/ldapadd -x -D "cn=Manager,dc=example,dc=com" -w admin0 -f /configs/init.ldif
/app/bin/ldapadd -x -D "cn=Manager,dc=example,dc=com" -w admin0 -f /configs/employee.ldif
/app/bin/ldapadd -x -D "cn=Manager,dc=example,dc=com" -w admin0 -f /configs/facility.ldif

/app/bin/ldapmodify -Q -Y EXTERNAL -H ldapi:/// -f /configs/add-configpw.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/cosine.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/inetorgperson.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi:/// -f /app/etc/openldap/schema/nis.ldif

/app/bin/ldapadd -x -D "cn=Manager,dc=example,dc=com" -w admin0 -f /configs/group-container.ldif
/app/bin/ldapadd -x -D "cn=Manager,dc=example,dc=com" -w admin0 -f /configs/user-groups.ldif
/app/bin/ldapadd -x -D "cn=Manager,dc=example,dc=com" -w admin0 -f /configs/users.ldif

while :
do
    sleep 10
done
