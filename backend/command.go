package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/OttyLab/openldap-shell/db"
)

func Add(parameter Parameter, db *db.Db) error {
	key, entry := toEntry(parameter)
	err := (*db).Write(key, entry)
	return err
}

func Search(parameter Parameter, db *db.Db) (string, error) {
	entries, err := (*db).Read()
	if err != nil {
		return "", err
	}

	scope := parameter["scope"][0]
	base := parameter["base"][0]
	filter := parameter["filter"][0]
	deref, _ := strconv.Atoi(parameter["deref"][0])
	sizeLimit, _ := strconv.Atoi(parameter["sizelimit"][0])

	low := "1"
	if scope == "0" {
		low = "0"
	}
	re, err := regexp.Compile("^([^=]+=[^=,]+,){" + low + "," + scope + "}" + base)
	if err != nil {
		return "", err
	}

	filtered := searchIntenal(entries, re, filter, sizeLimit, deref)

	return fromEntries(filtered), nil
}

func Compare(parameter Parameter, db *db.Db) string {
	entries, err := (*db).Read()
	if err != nil {
		return "RESULT\ncode: 34\n"
	}

	dn := parameter["dn"][0]
	if _, ok := entries[dn]; !ok {
		return "RESULT\ncode: 32\n"
	}

	target := entries[dn]
	result := false

	for key, values := range parameter {
		if key == "dn" || key == "msgid" || key == "suffix" {
			continue
		}

		if attributes, ok := target[key]; ok {
			for _, attribute := range attributes {
				if attribute == values[0] {
					result = true
				}
			}
		} else {
			return "RESULT\ncode: 32\n"
		}
	}

	if result {
		return "RESULT\ncode: 6\n"
	}

	return "RESULT\ncode: 5\n"
}

func searchIntenal(entries db.Entries, re *regexp.Regexp, filter string, sizeLimit int, deref int) db.Entries {
	filtered := make(db.Entries)
	aliases := make(db.Entries)

	cnt := 0
Loop:
	for dn, entry := range entries {
		if sizeLimit <= cnt {
			break
		}

		if re.Match([]byte(dn)) {
			if deref > 0 {
				for _, value := range entry["objectClass"] {
					if value == "alias" {
						aliasDn := entry["aliasedObjectName"][0]
						aliases[aliasDn] = entries[aliasDn]
						continue Loop
					}
				}
			}

			if Filter(filter, entry) {
				filtered[dn] = entry
				cnt++
			}
		}
	}

	for dn, entry := range aliases {
		if sizeLimit <= cnt {
			break
		}

		if Filter(filter, entry) {
			filtered[dn] = entry
			cnt++
		}
	}

	return filtered
}

// TODO: parse filter propery
func doesMatchFilter(entry db.Entry, filter string) bool {
	re := regexp.MustCompile("[^!]{0,1}\\(([^(&=]+?)=(.+?)\\)")
	matches := re.FindAllStringSubmatch(filter, -1)
	if matches == nil {
		return false
	}

	ret := false

	// TODO: Hack (phpLDAPadmin queries with one `objectClass=*`)
	if len(matches) == 1 && matches[0][1] == "objectClass" {
		return true
	}

	for _, match := range matches {
		re = regexp.MustCompile("(?i)" + strings.Replace(match[2], "*", ".*", -1))
		// TODO: Hack (apache uses `objectClass=*` first which wants to skip)
		if match[1] == "objectClass" {
			continue
		}

		if attributes, ok := entry[match[1]]; ok {
			for _, attribute := range attributes {
				ret = ret || re.MatchString(attribute)
			}
		}

		// TODO: Hack (sshd uses `objectClass=*` second which want to ignore)
		return ret
	}

	return ret
}
