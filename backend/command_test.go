package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"testing"

	"github.com/OttyLab/openldap-shell/db"
)

func TestAdd(t *testing.T) {
	inbuf := bytes.NewBufferString("{}")
	outbuf := bytes.NewBuffer(make([]byte, 0))
	driver := db.Db(db.NewJsonDB(inbuf, outbuf))

	parameter := Parameter{"dn": []string{"foo"}, "uid": []string{"bar"}}
	err := Add(parameter, &driver)

	if err != nil {
		t.Error(err)
	}

	scanner := bufio.NewScanner(outbuf)
	for scanner.Scan() {
		entries := new(db.Entries)
		err := json.Unmarshal(scanner.Bytes(), entries)
		if err != nil {
			t.Error(err)
		}
		if _, ok := (*entries)["foo"]; !ok {
			t.Fail()
		}
		if (*entries)["foo"]["uid"][0] != "bar" {
			t.Fail()
		}
	}
}
func TestSearch1(t *testing.T) {
	inbuf := bytes.NewBufferString(`
	{
		"cn=taro.yamada,ou=Employee,dc=example,dc=com":{
			"dn": ["cn=taro.yamada,ou=Employee,dc=example,dc=com"],
			"uid": ["taro"]
		},
		"dc=example,dc=com":{
			"dn": ["dc=example,dc=com"],
			"uid": ["taro"]
		}
	}
	`)

	outbuf := bytes.NewBuffer(make([]byte, 0))
	driver := db.Db(db.NewJsonDB(inbuf, outbuf))

	parameter := Parameter{
		"base":      []string{"dc=example,dc=com"},
		"scope":     []string{"2"},
		"sizelimit": []string{"500"},
		"filter":    []string{"(uid=*taro*)"},
	}
	result, _ := Search(parameter, &driver)

	expected := `dn: cn=taro.yamada,ou=Employee,dc=example,dc=com
uid: taro

RESULT
code: 0
`

	if result != expected {
		t.Error("does not match")
		println(result)
	}
}

func TestSearch2(t *testing.T) {
	inbuf := bytes.NewBufferString(`
	{
		"cn=taro.yamada,ou=Employee,dc=example,dc=com":{
			"dn": ["cn=taro.yamada,ou=Employee,dc=example,dc=com"],
			"objectClass": ["foo", "bar"],
			"cn": ["Taro Yamada"]
		},
		"cn=jiro.sato,ou=Employee,dc=example,dc=com":{
			"dn": ["cn=jiro.sato,ou=Employee,dc=example,dc=com"],
			"objectClass": ["bar"],
			"cn": ["Jiro Sato"]
		}
	}
	`)

	outbuf := bytes.NewBuffer(make([]byte, 0))
	driver := db.Db(db.NewJsonDB(inbuf, outbuf))

	parameter := Parameter{
		"base":      []string{"dc=example,dc=com"},
		"scope":     []string{"2"},
		"sizelimit": []string{"500"},
		"filter":    []string{" (&(|(cn=*taro*)(givenName=*taro*)(sn=*taro*)(?mozillaNickname=*taro*)(mail=*taro*)(?mozillaSecondEmail=*taro*)(&(description=*taro*))(o=*taro*)(ou=*taro*)(title=*taro*)(?mozillaWorkUrl=*taro*)(?mozillaHomeUrl=*taro*)))"},
		//"filter": []string{"foo=bar"},
	}

	result, _ := Search(parameter, &driver)

	expected := `dn: cn=taro.yamada,ou=Employee,dc=example,dc=com
objectClass: foo
objectClass: bar
cn: Taro Yamada

RESULT
code: 0
`

	if result != expected {
		t.Error("does not match")
		println(result)
	}
}
