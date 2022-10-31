#!/bin/sh

PORT=3389

/app/sbin/slaptest -f /data/slapd.conf.shell -F /app/etc/openldap/slapd.d
sleep 1

/app/libexec/slapd -h "ldap://0.0.0.0:${PORT}/ ldapi://0.0.0.0:${PORT}/"

/app/bin/ldapadd -H ldapi://0.0.0.0:${PORT}/ -x -D "cn=Manager,dc=example" -w admin0 -f /data/init.ldif
sleep 1

/app/bin/ldapmodify -Q -Y EXTERNAL -H ldapi://0.0.0.0:${PORT}/ -f /configs/add-configpw.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi://0.0.0.0:${PORT}/ -f /app/etc/openldap/schema/cosine.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi://0.0.0.0:${PORT}/ -f /app/etc/openldap/schema/inetorgperson.ldif
/app/bin/ldapadd -Q -Y EXTERNAL -H ldapi://0.0.0.0:${PORT}/ -f /app/etc/openldap/schema/nis.ldif


while :
do
    sleep 10
done
