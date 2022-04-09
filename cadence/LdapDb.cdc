pub contract LdapDb {
    pub let EntryStoragePath: StoragePath
    pub let EntryPublicPath: PublicPath

    pub let entries: {Address: Capability<&Entry{Reader}>}

    pub resource interface Reader {
        pub var attributes: {String: [String]} 
    }

    pub resource Entry: Reader {
        pub var attributes: {String: [String]} 

        init() {
            self.attributes = {}
        }

        pub fun setAttributes(attributes: {String: [String]}) {
            self.attributes = attributes
        }
    }

    pub fun createEntry(): @Entry {
        return <-create Entry()
    }

    pub fun setCapability(address: Address, capability: Capability<&Entry{Reader}>) {
        self.entries[address] = capability
    }

    init() {
        self.EntryStoragePath = /storage/LdapDbEntry
        self.EntryPublicPath = /public/LdapDbEntry
        self.entries = {}
    }
}
 