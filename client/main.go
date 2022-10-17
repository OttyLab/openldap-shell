package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/templates"
	"google.golang.org/grpc"
)

const NODE = "127.0.0.1:3569"
const SIGN_ALGO = "ECDSA_P256"
const HASH_ALGO = "SHA3_256"

const SERVICE_ADDRESS = "f8d6e0586b0a20c7"
const SERVICE_SIGN_ALGO = "ECDSA_P256"
const SERVICE_SK = "b53b3a5c476d4ef878601a2b5b5eb629f4885b3ab43395a81eb57ade2edbd99d"

type Attribute []string
type Entry map[string]Attribute
type Entries map[string]Entry

func generateKeys() crypto.PrivateKey {
	seed := make([]byte, crypto.MinSeedLength)
	_, err := rand.Read(seed)
	if err != nil {
		panic(err)
	}

	sk, err := crypto.GeneratePrivateKey(crypto.StringToSignatureAlgorithm(SIGN_ALGO), seed)
	if err != nil {
		panic(err)
	}

	return sk
}

func createAccount(client *client.Client, sk *crypto.PrivateKey) flow.Identifier {
	signAlgo := crypto.StringToSignatureAlgorithm(SIGN_ALGO)
	hashAlgo := crypto.StringToHashAlgorithm(HASH_ALGO)
	accountKey := flow.NewAccountKey().
		SetPublicKey((*sk).PublicKey()).
		SetSigAlgo(signAlgo).
		SetHashAlgo(hashAlgo).
		SetWeight(flow.AccountKeyWeightThreshold)

	ctx := context.Background()
	serviceAddress := flow.HexToAddress(SERVICE_ADDRESS)
	serviceAccount, err := client.GetAccountAtLatestBlock(ctx, serviceAddress)
	if err != nil {
		panic(err)
	}

	serviceAccountKey := serviceAccount.Keys[0]
	blockHeader, err := client.GetLatestBlockHeader(ctx, true)
	if err != nil {
		panic(err)
	}

	tx := templates.CreateAccount([]*flow.AccountKey{accountKey}, nil, serviceAddress).
		SetProposalKey(serviceAddress, serviceAccountKey.Index, serviceAccountKey.SequenceNumber).
		SetPayer(serviceAddress).
		SetGasLimit(uint64(100)).
		SetReferenceBlockID(blockHeader.ID) //added

	serviceSignAlgo := crypto.StringToSignatureAlgorithm(SERVICE_SIGN_ALGO)
	serviceSk, err := crypto.DecodePrivateKeyHex(serviceSignAlgo, SERVICE_SK)
	if err != nil {
		panic(err)
	}

	serviceSigner := crypto.NewInMemorySigner(serviceSk, serviceAccountKey.HashAlgo)
	err = tx.SignEnvelope(serviceAddress, serviceAccountKey.Index, serviceSigner)
	if err != nil {
		panic(err)
	}

	err = client.SendTransaction(ctx, *tx)
	if err != nil {
		panic(err)
	}

	return tx.ID()
}

func getAddress(client *client.Client, txID flow.Identifier) flow.Address {
	ctx := context.Background()

	result, err := client.GetTransactionResult(ctx, txID)
	if err != nil {
		panic("failed to get tx result")
	}

	var address flow.Address

	if result.Status == flow.TransactionStatusSealed {
		for _, event := range result.Events {
			if event.Type == flow.EventAccountCreated {
				accountCreatedEvent := flow.AccountCreatedEvent(event)
				address = accountCreatedEvent.Address()
			}
		}
	}

	return address
}

func getTransactionResult(client *client.Client, txID flow.Identifier) bool {
	ctx := context.Background()

	result, err := client.GetTransactionResult(ctx, txID)
	if err != nil {
		panic("failed to get tx result")
	}

	return result.Status == flow.TransactionStatusSealed
}

func getCreateEntryScript() string {
	script := `
		import LdapDb from 0xf8d6e0586b0a20c7

		transaction {
		    prepare(acct: AuthAccount) {
		        let entry <- LdapDb.createEntry()
		        acct.save(<-entry, to: LdapDb.EntryStoragePath)
		        let capability = acct.link<&LdapDb.Entry{LdapDb.Reader}>(LdapDb.EntryPublicPath, target: LdapDb.EntryStoragePath) ?? panic("failed capability")
		        LdapDb.setCapability(address: acct.address, capability: capability)
		        log("tx done")
		    }
		}
	`

	return script
}

func getSetAttriutesScript(entry Entry) string {
	script := `
		import LdapDb from 0xf8d6e0586b0a20c7

		transaction {
   			prepare(acct: AuthAccount) {
       			let entryRef = acct.borrow<&LdapDb.Entry>(from: LdapDb.EntryStoragePath) ?? panic("failed to borrow")
       			entryRef.setAttributes(attributes: $ATTRIBUTES$)
       			log("setAttributes done")
   			}
		}
	`

	bs, err := json.Marshal(entry)
	if err != nil {
		panic(err)
	}

	return strings.Replace(script, "$ATTRIBUTES$", string(bs), 1)
}

func sendTransaction(client *client.Client, address *flow.Address, sk *crypto.PrivateKey, script string) flow.Identifier {
	ctx := context.Background()

	account, err := client.GetAccountAtLatestBlock(ctx, *address)
	if err != nil {
		panic(err)
	}

	accountKey := account.Keys[0]

	blockHeader, err := client.GetLatestBlockHeader(ctx, true)
	tx := flow.NewTransaction().
		SetScript([]byte(script)).
		SetGasLimit(100).
		SetProposalKey(*address, 0, accountKey.SequenceNumber).
		SetReferenceBlockID(blockHeader.ID).
		SetPayer(account.Address).
		AddAuthorizer(account.Address)

	signer := crypto.NewInMemorySigner(*sk, crypto.StringToHashAlgorithm(HASH_ALGO))
	err = tx.SignEnvelope(account.Address, 0, signer)
	if err != nil {
		panic(err)
	}

	err = client.SendTransaction(ctx, *tx)
	if err != nil {
		panic(err)
	}

	return tx.ID()
}

func createClient() (*client.Client, error) {
	return client.New(NODE, grpc.WithInsecure())
}

func createAccountAndEntry(client *client.Client, entry Entry) {
	sk := generateKeys()
	txID := createAccount(client, &sk)
	fmt.Println("create account tx id=" + txID.String())

	blockTime := 5 * time.Second
	time.Sleep(blockTime)

	address := getAddress(client, txID)
	fmt.Println("created address=" + address.Hex())

	txID = sendTransaction(client, &address, &sk, getCreateEntryScript())
	fmt.Println("create entry tx id=" + txID.String())

	time.Sleep(blockTime)

	fmt.Println("create entry tx result=" + strconv.FormatBool(getTransactionResult(client, txID)))

	txID = sendTransaction(client, &address, &sk, getSetAttriutesScript(entry))
	fmt.Println("set attribute entry tx id=" + txID.String())

	time.Sleep(blockTime)

	fmt.Println("set attribute tx result=" + strconv.FormatBool(getTransactionResult(client, txID)))

}

func main() {
	if len(os.Args) != 2 {
		println("need db file parameter")
		os.Exit(1)
	}

	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		println("cannot open db file")
		os.Exit(1)
	}

	entries := new(Entries)
	err = json.Unmarshal(bs, entries)
	if err != nil {
		println("failed to parse db file")
		os.Exit(1)
	}

	client, err := createClient()
	if err != nil {
		println("failed to initialize client")
		os.Exit(1)
	}

	for key, entry := range *entries {
		fmt.Println("== creating " + key + " user ==")
		createAccountAndEntry(client, entry)
	}
}
