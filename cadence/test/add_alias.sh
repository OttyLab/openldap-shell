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
                        "value": "cn=taro.yamada,dc=org0,dc=example"
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
                                        {"type":"String", "value": "cn=taro.yamada,dc=org0,dc=example"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "cn"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "taro.yamada"}
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
                        "value": "cn=jiro.sato,dc=org1,dc=example"
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
                                        {"type":"String", "value": "cn=jiro.sato,dc=org1,dc=example"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "cn"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "jiro.sato"}
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
                                        {"type":"String", "value": "cn=jiro.sato,dc=person,dc=example"}
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