package main

import (
	"flag"
	"fmt"
	"github.com/ohze/flash-socket-policy-docker/fsp"
	"log"
	"os"
)

var (
	VERSION = "1.0.0"
)

var version = flag.Bool("version", false, "show gofsp version")
var filename = flag.String("file", "", "policy file")

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("gofsp - %s\n", VERSION)
		os.Exit(0)
	}

	s := fsp.NewServer()

	// load policy from file
	if len(*filename) > 0 {
		file, err := os.Open(*filename)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		s.LoadPolicy(file)
	}

	s.ListenAndServe()
}
