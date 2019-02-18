package internals

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func checkData(data []string) int {

	if len(data) == 1 || len(data) > 2 {
		command := data[0]
		if strings.EqualFold(command, createAParkingLot) {
			log.Printf("Incomplete/Invalid %s command with %s data, try i.e %s num \n", command, data, createAParkingLot)
			return -1
		}
		if strings.EqualFold(command, freeSlot) {
			log.Printf("Incomplete/Invalid %s command with %s data, try i.e %s solNum \n", command, data, command)
			return -1
		}

		if strings.EqualFold(command, regNumbersWithColor) {
			log.Printf("Incomplete/Invalid %s command with %s data, try i.e %s colorname \n", command, data, regNumbersWithColor)
			return -1
		}

		if strings.EqualFold(command, slotNumbersWithColor) {
			log.Printf("Incomplete/Invalid %s command with %s data, try i.e %s slotNum \n", command, data, slotNumbersWithColor)
			return -1
		}

		if strings.EqualFold(command, slotNumberWithReg) {
			log.Printf("Incomplete/Invalid %s command with %s data, try i.e %s regNum \n", command, data, slotNumberWithReg)
			return -1
		}
	}
	return 1
}

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
		if len(data) == 0 {
			log.Printf("Command not found \n")
			continue
		}
		if checkData(data) != -1 {
			processCommand(data)
		}
	}
}

func ReadFromConsole() {

	fmt.Println("Reading input from console ...")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")
		if len(data) == 0 {
			log.Printf("Command not found \n")
			continue
		}
		if checkData(data) != -1 {
			processCommand(data)
		}
		if line == "exit" {
			os.Exit(0)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
