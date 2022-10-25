package main

import (
	"testing"

	"github.com/OttyLab/openldap-shell/db"
)

var entries []db.Entry = []db.Entry{
	{
		"dn":            db.Attribute{"cn=taro.yamada,ou=Employee,dc=example,dc=com"},
		"cn":            db.Attribute{"taro.yamada"},
		"mail":          db.Attribute{"taro.yamada@example.com"},
		":userPassword": db.Attribute{"1234"},
		"objectClass":   db.Attribute{"inetOrgPerson", "posixAccount"},
		"uid":           db.Attribute{"taro.yamada"},
		"sn":            db.Attribute{"Yamada"},
	},
	{
		"dn":            db.Attribute{"cn=jiro.sato,ou=Employee,dc=example,dc=com"},
		"cn":            db.Attribute{"jiro.sato"},
		"mail":          db.Attribute{"j-dato@example.com"},
		":userPassword": db.Attribute{"abcd"},
		"objectClass":   db.Attribute{"inetOrgPerson", "posixAccount"},
		"uid":           db.Attribute{"jiro.sato"},
		"sn":            db.Attribute{"Sato"},
	},
	{
		"dn":            db.Attribute{"cn=saburo.suzuki,ou=Employee,dc=example2,dc=com"},
		"cn":            db.Attribute{"saburo.suzuki"},
		"mail":          db.Attribute{"saburo@example2.com"},
		":userPassword": db.Attribute{"a1b2"},
		"objectClass":   db.Attribute{"inetOrgPerson"},
		"uid":           db.Attribute{"saburo"},
		"sn":            db.Attribute{"Suzuki"},
	},
	{
		"dn":            db.Attribute{"cn=shiro.yamada,ou=Employee,dc=example2,dc=com"},
		"cn":            db.Attribute{"shiro.yamada"},
		"mail":          db.Attribute{"s-yamada@example2.com"},
		":userPassword": db.Attribute{"c3d4"},
		"objectClass":   db.Attribute{"posixAccount"},
		"uid":           db.Attribute{"s-yamada"},
		"sn":            db.Attribute{"Yamada"},
	},
}

func TestFilter1(t *testing.T) {
	filter := "(&(sn=Yamada)(mail=*example*))"

	result := Filter(filter, entries[0])
	if !result {
		t.Error("error with dn=" + entries[0]["dn"][0])
	}

	result = Filter(filter, entries[1])
	if result {
		t.Error("error with dn=" + entries[0]["dn"][1])
	}

	result = Filter(filter, entries[2])
	if result {
		t.Error("error with dn=" + entries[0]["dn"][2])
	}

	result = Filter(filter, entries[3])
	if !result {
		t.Error("error with dn=" + entries[0]["dn"][3])
	}
}

func TestFilter2(t *testing.T) {
	filter := "(|(sn=Yamada)(mail=*example.com))"

	result := Filter(filter, entries[0])
	if !result {
		t.Error("error with dn=" + entries[0]["dn"][0])
	}

	result = Filter(filter, entries[1])
	if !result {
		t.Error("error with dn=" + entries[0]["dn"][1])
	}

	result = Filter(filter, entries[2])
	if result {
		t.Error("error with dn=" + entries[0]["dn"][2])
	}

	result = Filter(filter, entries[3])
	if !result {
		t.Error("error with dn=" + entries[0]["dn"][3])
	}
}

func TestFilter3(t *testing.T) {
	filter := "(&(!(sn=Yamada))(mail=*example*))"

	result := Filter(filter, entries[0])
	if result {
		t.Error("error with dn=" + entries[0]["dn"][0])
	}

	result = Filter(filter, entries[1])
	if !result {
		t.Error("error with dn=" + entries[0]["dn"][1])
	}

	result = Filter(filter, entries[2])
	if !result {
		t.Error("error with dn=" + entries[0]["dn"][2])
	}

	result = Filter(filter, entries[3])
	if result {
		t.Error("error with dn=" + entries[0]["dn"][3])
	}
}

func TestFilter4(t *testing.T) {
	filter := "(&(|(cn=*taro*)(givenName=*taro*)(sn=*taro*)(?mozillaNickname=*taro*)(mail=*taro*)(?mozillaSecondEmail=*taro*)(&(description=*taro*))(o=*taro*)(ou=*taro*)(title=*taro*)(?mozillaWorkUrl=*taro*)(?mozillaHomeUrl=*taro*)))"

	result := Filter(filter, entries[0])
	if !result {
		t.Error("error with dn=" + entries[0]["dn"][0])
	}

	result = Filter(filter, entries[1])
	if result {
		t.Error("error with dn=" + entries[0]["dn"][1])
	}

	result = Filter(filter, entries[2])
	if result {
		t.Error("error with dn=" + entries[0]["dn"][2])
	}

	result = Filter(filter, entries[3])
	if result {
		t.Error("error with dn=" + entries[0]["dn"][3])
	}
}
