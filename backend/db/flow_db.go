package db

import (
	"context"
	"fmt"
	"os"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

type FlowDb struct {
	node     string
	contract string
	address  string
	client   *client.Client
}

func NewFlowDb(node string, contract string, address string) *FlowDb {
	ret := new(FlowDb)
	ret.node = node
	ret.contract = contract
	ret.address = address

	clinet, err := client.New(node, grpc.WithInsecure())
	if err != nil {
		println("failed to initialize client")
		os.Exit(1)
	}

	ret.client = clinet

	return ret
}

func (flowDb *FlowDb) Write(dn string, entry Entry) error {
	return fmt.Errorf("Flow DB does not support Write")
}

func (flowDb *FlowDb) Read() (Entries, error) {
	return flowDb.read()
}

func (flowDb *FlowDb) read() (Entries, error) {
	value, err := flowDb.getEntriesFromFlow()
	if err != nil {
		return nil, err
	}

	entries := make(Entries)
	for _, entry_pair := range value.(cadence.Dictionary).Pairs {
		entry := make(Entry)
		dn := ""
		for _, attribute_pair := range entry_pair.Value.(cadence.Dictionary).Pairs {
			attributes := make(Attribute, 0)

			for _, value := range attribute_pair.Value.(cadence.Array).Values {
				attributes = append(attributes, value.ToGoValue().(string))
			}

			key := attribute_pair.Key.ToGoValue().(string)
			entry[key] = attributes
			if key == "dn" {
				dn = attributes[0]
			}
		}
		entries[dn] = entry
	}

	return entries, nil
}

func (flowDb *FlowDb) getEntriesFromFlow() (cadence.Value, error) {
	ctx := context.Background()

	// TODO:
	script := `
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
	`

	return flowDb.client.ExecuteScriptAtLatestBlock(ctx, []byte(script), nil)
}
