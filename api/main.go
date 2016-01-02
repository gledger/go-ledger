package main

import (
	"flag"
	"os"

	"github.com/zombor/go-ledger"
	"github.com/zombor/go-ledger/Godeps/_workspace/src/github.com/dougblack/sleepy"
)

func main() {
	listen := flag.Int("listen", 0,
		"TCP address (host:port) on which to listen for HTTP connections."+
			" Defaults to a random port."+
			" See http://golang.org/pkg/net/#Dial for examples.")
	journalPath := flag.String("journal", "", "File path to ledger journal file. Required")

	flag.Parse()

	if *journalPath == "" {
		println("-journal is a required flag")
		os.Exit(-1)
	}

	if *listen < 1 {
		println("-listen is a required flag")
		os.Exit(-1)
	}

	journal := ledger.NewJournal(
		ledger.NewFileReader(*journalPath),
	)

	rootHandler := rootHandler{journal: journal}
	transactionsHandler := transactionsHandler{
		journalReader: journal,
		journalWriter: journal,
	}

	api := sleepy.NewAPI()
	api.AddResource(rootHandler, "/")
	api.AddResource(transactionsHandler, "/transactions")
	api.Start(*listen)
}
