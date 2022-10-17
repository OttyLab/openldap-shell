package db

import "testing"

func TestRead(t *testing.T) {
	flowDb := NewFlowDb("127.0.0.1:3569", "LdapDb", "0xf8d6e0586b0a20c7")
	flowDb.Read()
}
