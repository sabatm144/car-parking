package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"parking-lot/internals"
)

func main() {

	mode := flag.Int("mode", 1, "Select a mode to run")
	filePath := flag.String("filePath", "", "Input file path please!")

	flag.Parse()

	// get the actual option values.
	fmt.Println("mode:", *mode)
	fmt.Println("filePath:", *filePath)

	if *mode == 2 {
		fmt.Println("Processing file, please wait")
		if *filePath == "" {
			fmt.Println("Hey you didn’t mention your input file path")
			flag.PrintDefaults()
			os.Exit(1)
		}

		internals.ReadFromFile(*filePath)
		os.Exit(1)
	}

	fmt.Println("Interactive mode: please type your operations")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		//TODO: execute the command

		if text == "exit" {
			fmt.Println("Good bye")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
