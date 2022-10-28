import LdapDb from 0xf8d6e0586b0a20c7

transaction(entry: LdapDb.Entry) {
    prepare(acct: AuthAccount) {
        let ditRef = acct.borrow<&LdapDb.Dit>(from: LdapDb.DitStoragePath) ?? panic("failed to borrow")
        ditRef.setEntry(entry: entry)
        log("setEntry done")
    }
}
 