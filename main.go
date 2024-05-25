package main

import (
	"Math-Skills/readandprocess"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error. Incorrect number of arguments passed. Usage: go run . data.txt")
		os.Exit(1)
	}

	arguments := os.Args[1:]

	if filepath.Ext(arguments[0]) != ".txt" {
		fmt.Println("Error. Incorrect file extension used!")
		os.Exit(1)
	}

	if arguments[0] != "data.txt" {
		fmt.Println("Error. File to access must be 'data.txt'")
		os.Exit(1)
	}

	statisticsResult := readandprocess.ReadAndProcess(arguments[0])

	fmt.Println(statisticsResult)
}
