package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/OttyLab/openldap-shell/db"
)

type Command int
type Parameter map[string][]string

const (
	ADD Command = iota
	SEARCH
	COMPARE
	BIND
	UNBIND
)

func parse(scanner *bufio.Scanner) (Command, Parameter, error) {
	if !scanner.Scan() {
		return -1, nil, fmt.Errorf("command not found")
	}

	switch scanner.Text() {
	case "ADD":
		return ADD, readParams(scanner), nil
	case "SEARCH":
		return SEARCH, readParams(scanner), nil
	case "COMPARE":
		return COMPARE, readParams(scanner), nil
	case "BIND":
		return BIND, readParams(scanner), nil
	case "UNBIND":
		return UNBIND, readParams(scanner), nil
	default:
		return -1, nil, fmt.Errorf("unexpected command")
	}
}

func readParams(scanner *bufio.Scanner) Parameter {
	result := make(Parameter)
	re := regexp.MustCompile(`([^:]+) *:+ *(.+)`)
	reBase64 := regexp.MustCompile(`([^:]+) *:: *(.+)`)

	for scanner.Scan() {
		attribute := scanner.Text()
		r := re.FindStringSubmatch(attribute)
		if r == nil {
			continue
		}

		// Note
		// `::` delimiter means the value is Base64 encoded.
		// Here `:` is added in the head of attribute name to show the value is Base64.
		// According to the [specification](https://www.rfc-editor.org/rfc/rfc4512.txt),
		// `:` is not used as attribute names and it's safe.
		//
		var attributeName string
		if reBase64.MatchString(attribute) {
			attributeName = ":" + r[1]
		} else {
			attributeName = r[1]
		}

		if _, ok := result[attributeName]; !ok {
			result[attributeName] = make([]string, 0)
		}

		result[attributeName] = append(result[r[1]], r[2])
	}

	return result
}

func toEntry(parameter Parameter) (dn string, entry db.Entry) {
	dn = parameter["dn"][0]
	entry = make(db.Entry)

	for key, values := range parameter {
		var attributes db.Attribute
		for _, value := range values {
			attributes = append(attributes, value)
		}

		entry[key] = attributes
	}

	return dn, entry
}

func fromEntries(entries db.Entries) string {
	var sb strings.Builder
	for _, entry := range entries {
		for attributeName, attributes := range entry {
			for _, attribute := range attributes {
				if attributeName[0] == ':' {
					sb.WriteString(attributeName[1:] + ":: " + attribute + "\n")
				} else {
					sb.WriteString(attributeName + ": " + attribute + "\n")
				}
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString("RESULT\n")
	sb.WriteString("code: 0\n")

	return sb.String()
}
