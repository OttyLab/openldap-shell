#!/bin/sh

for i in $(seq ${1} ${2})
do
    echo "dn: uid=john_${i},ou=Employee,dc=example,dc=com"
    echo "objectClass: inetOrgPerson"
    echo "uid: john_${i}"
    echo "cn: John_${i} Doe_${i}"
    echo "sn: Doe_${i}"
    echo "givenname: John_${i}"
    echo "mail: john_${i}.doe_${i}@example.com"
    echo "ou: System development"
    echo "o: Capybara Corp"
    echo "title: Engineer"
    echo "streetaddress: Shinagawa-ku, Tokyo"
    echo "telephonenumber: 1234567"
    echo ""
done
