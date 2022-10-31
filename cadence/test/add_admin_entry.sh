flow transactions send ./transactions/set_entry.cdc --signer admin --args-json '[
    {
        "type": "Struct",
        "value": {
            "id": "A.f8d6e0586b0a20c7.LdapDb.Entry",
            "fields": [
                {
                    "name": "dn",
                    "value": {
                        "type": "String",
                        "value": "dc=example"
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
                                        {"type":"String", "value": "dc=example"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "dc"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "dc=example"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "o"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "Example"}
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
                                        {"type":"String", "value": "organization"},
                                        {"type":"String", "value": "dcObject"}
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

flow transactions send ./transactions/set_entry.cdc --signer admin --args-json '[
    {
        "type": "Struct",
        "value": {
            "id": "A.f8d6e0586b0a20c7.LdapDb.Entry",
            "fields": [
                {
                    "name": "dn",
                    "value": {
                        "type": "String",
                        "value": "cn=Manager,dc=example"
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
                                        {"type":"String", "value": "cn=Manager,dc=example"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "cn"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "Manager"}
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
                                        {"type":"String", "value": "organizationalRole"}
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

flow transactions send ./transactions/set_entry.cdc --signer admin --args-json '[
    {
        "type": "Struct",
        "value": {
            "id": "A.f8d6e0586b0a20c7.LdapDb.Entry",
            "fields": [
                {
                    "name": "dn",
                    "value": {
                        "type": "String",
                        "value": "dc=person,dc=example"
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
                                        {"type":"String", "value": "dc=person,dc=example"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "dc"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "dc=person"}
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