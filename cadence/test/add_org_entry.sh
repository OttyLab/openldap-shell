flow transactions send ./transactions/set_entry.cdc --signer org0 --args-json '[
    {
        "type": "Struct",
        "value": {
            "id": "A.f8d6e0586b0a20c7.LdapDb.Entry",
            "fields": [
                {
                    "name": "dn",
                    "value": {
                        "type": "String",
                        "value": "ou=org0,dc=example,dc=com"
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
                                        {"type":"String", "value": "ou=org0,dc=example,dc=com"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "ou"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "ou=org0"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "suffix"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "dc=example,dc=com"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "objectClass"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "organizationalUnit"}
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

flow transactions send ./transactions/set_entry.cdc --signer org1 --args-json '[
    {
        "type": "Struct",
        "value": {
            "id": "A.f8d6e0586b0a20c7.LdapDb.Entry",
            "fields": [
                {
                    "name": "dn",
                    "value": {
                        "type": "String",
                        "value": "ou=org1,dc=example,dc=com"
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
                                        {"type":"String", "value": "ou=org1,dc=example,dc=com"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "ou"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "ou=org1"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "suffix"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "dc=example,dc=com"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "objectClass"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "organizationalUnit"}
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