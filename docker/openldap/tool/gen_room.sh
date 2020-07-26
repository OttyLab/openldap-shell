#!/bin/sh

for i in $(seq ${1} ${2})
do
    echo "dn: commonName=room_${i},ou=Facility,dc=example,dc=com"
    echo "objectClass: room"
    echo "commonName: room_${i}"
    echo "roomNumber: ${i}"
    echo "description: meeting room"
    echo ""
done
