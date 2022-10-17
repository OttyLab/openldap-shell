import LdapDb from 0xf8d6e0586b0a20c7

transaction {
    prepare(acct: AuthAccount) {
        let entry <- LdapDb.createEntry()
        acct.save(<-entry, to: LdapDb.EntryStoragePath)
        let capability = acct.link<&LdapDb.Entry{LdapDb.Reader}>(LdapDb.EntryPublicPath, target: LdapDb.EntryStoragePath) ?? panic("failed capability")
        LdapDb.setCapability(address: acct.address, capability: capability)
        log("tx done")
    }
}
