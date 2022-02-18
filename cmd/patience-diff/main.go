package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	diff "github.com/thepudds/patience-diff"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "patience-diff file1 file2")
	}
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(2)
	}

	old, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	new, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	result := diff.Diff(os.Args[1], old, os.Args[2], new)
	os.Stdout.Write(result)
}
