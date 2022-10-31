flow transactions send ./transactions/set_entry.cdc --signer user0 --args-json '[
    {
        "type": "Struct",
        "value": {
            "id": "A.f8d6e0586b0a20c7.LdapDb.Entry",
            "fields": [
                {
                    "name": "dn",
                    "value": {
                        "type": "String",
                        "value": "cn=guest,dc=org1,dc=example"
                    }
                },
                {
                    "name": "attributes",
                    "value": {
                        "type": "Dictionary",
                        "value": [
                            {
                                "key": {"type": "String", "value": "dn"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "cn=guest,dc=org1,dc=example"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "cn"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "cn=guest"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "suffix"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "dc=example"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "objectClass"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "alias"},
                                        {"type":"String", "value": "extensibleObject"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "aliasedObjectName"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "cn=taro.yamada,dc=person,dc=example"}
                                    ]
                                }
                            }
                        ]
                    }
                }
            ]
        }
    }
]'

flow scripts execute ./scripts/get_entries.cdc