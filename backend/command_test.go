package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"strings"
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

	//
	//Expected:
	//dn: cn=taro.yamada,ou=Employee,dc=example,dc=com
	//uid: taro
	//
	//RESULT
	//code: 0
	//

	if strings.Index(result, "dn: cn=taro.yamada,ou=Employee,dc=example,dc=com") == -1 {
		t.Error("dn does not exist")
	}
	if strings.Index(result, "uid: taro") == -1 {
		t.Error("uid does not exist")
	}
	if strings.Index(result, "code: 0") == -1 {
		t.Error("code: 0 does not exist")
	}
	if strings.Index(result, "code: 0") == -1 {
		t.Error("code: 0 does not exist")
	}
}

func TestSearch2(t *testing.T) {
	inbuf := bytes.NewBufferString(`
	{
		"cn=taro.yamada,ou=Employee,dc=example,dc=com":{
			"dn": ["cn=taro.yamada,ou=Employee,dc=example,dc=com"],
			"objectClass": ["foo", "bar"],
			"cn": ["Taro Yamada"],
			"mail": ["taro.yamada@example.com"]
		},
		"cn=jiro.sato,ou=Employee,dc=example,dc=com":{
			"dn": ["cn=jiro.sato,ou=Employee,dc=example,dc=com"],
			"objectClass": ["bar"],
			"cn": ["Jiro Sato"],
			"mail": ["j-sato@example.com"]
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
	}

	result, _ := Search(parameter, &driver)

	//
	// Expected
	//dn: cn=taro.yamada,ou=Employee,dc=example,dc=com
	//objectClass: foo
	//objectClass: bar
	//cn: Taro Yamada
	//
	//RESULT
	//code: 0
	//
	if strings.Index(result, "dn: cn=taro.yamada,ou=Employee,dc=example,dc=com") == -1 {
		t.Error("dn does not exist")
	}
	if strings.Index(result, "objectClass: foo") == -1 {
		t.Error("objectClass foo does not exist")
	}
	if strings.Index(result, "objectClass: bar") == -1 {
		t.Error("objectClass bar does not exist")
	}
	if strings.Index(result, "cn: Taro Yamada") == -1 {
		t.Error("cn does not exist")
	}
	if strings.Index(result, "code: 0") == -1 {
		t.Error("code: 0 does not exist")
	}
	if strings.Index(result, "code: 0") == -1 {
		t.Error("code: 0 does not exist")
	}
}
