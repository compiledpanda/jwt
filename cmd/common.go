package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readStdIn() []byte {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Could not read from stdin", err)
		os.Exit(1)
	}
	return b
}

func readFile(file string) []byte {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Could not open file", file, err)
		os.Exit(1)
	}
	return b
}
