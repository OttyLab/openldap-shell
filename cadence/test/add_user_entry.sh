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
                        "value": "cn=taro.yamada,dc=person,dc=example"
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
                                        {"type":"String", "value": "cn=taro.yamada,dc=person,dc=example"}
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
                                        {"type":"String", "value": "dc=person"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "objectClass"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "inetOrgPerson"},
                                        {"type":"String", "value": "posixAccount"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "gidNumber"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "3000"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "uidNumber"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "3001"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "uid"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "taro.yamada"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": ":userPassword"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "e0NSWVBUfVRKVlJ5aVNNVUFRekE="}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "givenName"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "Taro"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "sn"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "Yamada"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "loginShell"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "/bin/bash"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "homeDirectory"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "/home/taro.yamada"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "mail"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "taro.yamada@example.com"}
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

flow transactions send ./transactions/set_entry.cdc --signer user1 --args-json '[
    {
        "type": "Struct",
        "value": {
            "id": "A.f8d6e0586b0a20c7.LdapDb.Entry",
            "fields": [
                {
                    "name": "dn",
                    "value": {
                        "type": "String",
                        "value": "cn=jiro.sato,dc=person,dc=example"
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
                                        {"type":"String", "value": "cn=jiro.sato,dc=person,dc=example"}
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
                                        {"type":"String", "value": "dc=person"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "objectClass"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "inetOrgPerson"},
                                        {"type":"String", "value": "posixAccount"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "gidNumber"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "3000"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "uidNumber"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "3002"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "uid"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "jiro.sato"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": ":userPassword"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "e0NSWVBUfUJLMnpVT2Vhc3NDYU0="}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "givenName"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "Jiro"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "sn"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "Sato"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "loginShell"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "/bin/bash"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "homeDirectory"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "/home/jiro.sato"}
                                    ]
                                }
                            },
                            {
                                "key": {"type": "String", "value": "mail"},
                                "value": {"type": "Array",
                                    "value": [
                                        {"type":"String", "value": "j-sato@example.com"}
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