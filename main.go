package main

import (
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
	// -mode = 2 for file input mode
	if *mode == 2 {
		fmt.Println("Processing file, please wait")
		fmt.Println("filePath:", *filePath)
		if *filePath == "" {
			fmt.Println("Hey you didn’t mention your input file path")
			flag.PrintDefaults()
			os.Exit(1)
		}

		internals.ReadFromFile(*filePath)
		os.Exit(1)
	}

	// -mode = 1/ with/without for interactive mode
	internals.ReadFromConsole()
}
