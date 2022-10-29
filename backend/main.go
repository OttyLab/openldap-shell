package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/OttyLab/openldap-shell/db"
)

func main() {
	f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		println("cannot open log file")
		os.Exit(1)
	}
	log.SetOutput(f)

	// JSON DB
	if _, err = os.Stat("db.json"); err != nil {
		if f, err := os.Create("db.json"); err != nil {
			log.Println(err)
			os.Exit(1)
		} else {
			f.WriteString("{}")
			f.Close()
		}
	}

	reader, err := os.Open("db.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer reader.Close()

	writer, err := os.OpenFile("db.json", os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer writer.Close()

	jsonDb := db.Db(db.NewJsonDB(reader, writer))

	// Flow DB
	flowDb := db.Db(db.NewFlowDb("flow-emulator:3569", "LdapDb", "0xf8d6e0586b0a20c7"))

	scanner := bufio.NewScanner(os.Stdin)
	command, parameter, err := parse(scanner)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//DEBUG
	//log.Println("[" + strconv.Itoa(int(command)) + "]")
	//for k, vs := range parameter {
	//	log.Printf(k + ": ")
	//	for _, v := range vs {
	//		log.Printf("%s, ", v)
	//	}
	//	log.Println("")
	//}

	switch command {
	case ADD:
		err := Add(parameter, &jsonDb)
		if err != nil {
			log.Println(err)
		}
	case SEARCH:
		resultJsonDb, err := Search(parameter, &jsonDb)
		if err != nil {
			log.Println(err)
		}

		resultFlowDb, err := Search(parameter, &flowDb)
		if err != nil {
			log.Println(err)
		}

		result := make(db.Entries)
		for key, value := range resultJsonDb {
			result[key] = value
		}
		for key, value := range resultFlowDb {
			result[key] = value
		}

		log.Println("[Search result in Flow DB]")
		log.Println(result)
		fmt.Print(fromEntries(result))
	case COMPARE:
		result := Compare(parameter, &jsonDb)
		if result != 6 {
			log.Println("[iCompare result in JSON DB]")
			log.Println(result)
			fmt.Print("RESULT\ncode: " + strconv.Itoa(result) + "\n")
			break
		}

		result = Compare(parameter, &flowDb)

		log.Println("[Compare result in Flow DB]")
		log.Println(result)
		fmt.Print("RESULT\ncode: " + strconv.Itoa(result) + "\n")
	case BIND:
		log.Println("[Bind result]")
	case UNBIND:
		log.Println("[Unbind result]")
	default:
		log.Println("TODO")
	}

	os.Exit(0)
}
