package internals

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFromFile(filePath string) {
	fmt.Println("Reading input from file ...")

	fp, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error in reading file %v \n", err)
		return
	}

	lines := strings.Split(string(fp), "\n")
	for _, line := range lines {
		data := strings.Split(line, " ")
		processCommand(data)
	}
}

func ReadFromConsole() {

	fmt.Println("Reading input from console ...")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")
		processCommand(data)
		if line == "exit" {
			os.Exit(0)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
