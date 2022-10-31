import LdapDb from 0xf8d6e0586b0a20c7

transaction (dn: String) {
    prepare(acct: AuthAccount) {
        let dit <- LdapDb.createDit(dn: dn)
        acct.save(<-dit, to: LdapDb.DitStoragePath)
        let capability = acct.link<&LdapDb.Dit{LdapDb.Reader}>(LdapDb.DitPublicPath, target: LdapDb.DitStoragePath) ?? panic("failed capability")
        LdapDb.setCapability(address: acct.address, capability: capability)
        log("tx done")
    }
}
