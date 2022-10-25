package main

import (
	"regexp"
	"strings"

	"github.com/OttyLab/openldap-shell/db"
)

func Filter(filter string, entry db.Entry) bool {
	_, result := dfs(filter, entry)
	return result
}

func dfs(substring string, entry db.Entry) (string, bool) {
	index := strings.Index(substring, "(")
	substring = substring[index+1:]

	switch h(substring) {
	case '&':
		sum := true
		result := false
		next := substring
		for {
			next, result = dfs(next, entry)
			sum = sum && result
			if next[0] == ')' {
				break
			}
		}
		return next[1:], sum
	case '|':
		sum := false
		result := false
		next := substring
		for {
			next, result = dfs(next, entry)
			sum = sum || result
			if next[0] == ')' {
				break
			}
		}
		return next[1:], sum
	case '!':
		next, result := dfs(substring, entry)
		return next[1:], !result
	default:
		return judge(substring, entry)
	}
}

func judge(s string, entry db.Entry) (string, bool) {
	index := strings.Index(s, ")")
	slice := strings.Split(s[0:index], "=")
	reg := "^" + strings.Replace(slice[1], "*", ".*", -1) + "$"

	result := false
	if vals, ok := entry[slice[0]]; ok {
		for _, val := range vals {
			if m, _ := regexp.MatchString(reg, val); m {
				result = true
				break
			}
		}
	}

	return s[index+1:], result
}

func h(s string) rune {
	return []rune(s)[0]
}
