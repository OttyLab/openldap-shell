import LdapDb from 0xf8d6e0586b0a20c7

transaction {
    prepare(acct: AuthAccount) {
        let entryRef = acct.borrow<&LdapDb.Entry>(from: LdapDb.EntryStoragePath) ?? panic("failed to borrow")
        entryRef.setAttributes(attributes: {"key1": ["value1-1", "value1-2"], "key2": ["value2-1", "value2-2"]})
        log("setAttributes done")
    }
}
 