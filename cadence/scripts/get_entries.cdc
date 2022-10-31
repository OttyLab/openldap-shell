import LdapDb from 0xf8d6e0586b0a20c7

pub fun main(): {String: {String: [String]}} {
    var entries: {String: {String: [String]}} = {}

    for key in LdapDb.dits.keys {
        if let capability = LdapDb.dits[key]{
            if let ditRef = capability.borrow() {
                log(key)
                log(ditRef.dn)
                for entry in ditRef.entries {
                    entries[entry.dn] = entry.attributes
                }
            }
        }
    }

    log(entries)
    return entries
}
 