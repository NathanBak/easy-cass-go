package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/NathanBak/easy-cass-go/pkg/easycass"
)

// The extractprops command line tool accepts an Astra secure connect bundle zip and will print out
// the following properties:
// - hostname
// - port
// - keyspace
// - certPEMBlock
// - keyPemBlock
// - pemCerts
//
// Using a "-q" flag will quote the values.  Using a "-j" flag will print them as JSON.

func main() {
	var secureBundlePath string
	quotes := false
	useJSON := false

	for i, arg := range os.Args {
		switch i {
		case 0: // program name
			continue
		case len(os.Args) - 1:
			secureBundlePath = arg
		default:
			if arg == "-q" {
				quotes = true
			} else if arg == "-j" {
				useJSON = true
			}
		}
	}

	props, err := easycass.ExtractProperties(secureBundlePath)
	if err != nil {
		log.Fatal(err)
	}

	if useJSON {
		buf, err := json.MarshalIndent(&props, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf))
		os.Exit(0)
	}

	format := "%s = %s\n"
	if quotes {
		format = "%s = \"%s\"\n"

	}

	for k, v := range props {
		fmt.Printf(format, k, v)
	}
}
