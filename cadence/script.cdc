import LdapDb from 0xf8d6e0586b0a20c7

pub fun main(): {String: {String: [String]}} {
    var entries: {String: {String: [String]}} = {}
    for key in LdapDb.entries.keys {
        if let capability = LdapDb.entries[key]{
            if let entryRef = capability.borrow() {
                entries[key.toString()] = entryRef.attributes
                log(key)
                log(entryRef.attributes)
            }
        }
    }

    return entries
}
 