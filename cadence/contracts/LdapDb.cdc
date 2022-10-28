pub contract LdapDb {
    pub let DitStoragePath: StoragePath
    pub let DitPublicPath: PublicPath

    pub let dits: {Address: Capability<&Dit{Reader}>}

    pub resource interface Reader {
        pub var dn: String
        pub var entries: [Entry]
    }

    pub struct Entry {
        pub var dn: String
        pub var attributes: {String: [String]} 

        init() {
            self.dn = ""
            self.attributes= {}
        }
    }

    pub resource Dit: Reader {
        pub var dn: String
        pub var entries: [Entry]

        init(dn: String) {
            self.dn = dn
            self.entries = []
        }

        pub fun setEntry(entry: Entry) {
            self.entries.append(entry)
        }
        //pub fun setAttributes(attributes: {String: [String]}) {
        //    self.attributes = attributes
        //}
    }

    pub fun createDit(dn: String): @Dit {
        return <-create Dit(dn: dn)
    }

    pub fun setCapability(address: Address, capability: Capability<&Dit{Reader}>) {
        self.dits[address] = capability
    }

    init() {
        self.DitStoragePath = /storage/LdapDbDit
        self.DitPublicPath = /public/LdapDbDit
        self.dits = {}
    }
}
 