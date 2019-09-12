package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

func stringToSeconds(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if i < 0 {
		fmt.Println("negative seconds not allowed")
		os.Exit(1)
	}
	return i
}
