package internals

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	createAParkingLot    = "create_parking_lot"
	parkVehicle          = "park"
	statusOfParkingSlots = "status"
	freeSlot             = "leave"
	regNumbersWithColor  = "registration_numbers_for_cars_with_colour"
	slotNumbersWithColor = "slot_numbers_for_cars_with_colour"
	slotNumberWithReg    = "slot_number_for_registration_number"
)

type vehicleInfo struct {
	Number  string
	Color   string
	Alloted bool
}

var (
	totalNumberOfParkingLots int
	parkingSlots             map[int]vehicleInfo
)

func checkData(dataType string, data []string) string {

	if len(data) == 0 {
		return fmt.Sprintf("Command not found")
	}

	if len(data) == 1 || len(data) > 2 {
		if dataType == createAParkingLot {
			return fmt.Sprintf("Incomplete/Invalid %s command with %s data, try i.e %s num", dataType, data, createAParkingLot)
		}
		if dataType == freeSlot {
			return fmt.Sprintf("Incomplete/Invalid %s command, try i.e %s solNum", data, dataType)
		}

		if dataType == regNumbersWithColor {
			return fmt.Sprintf("Incomplete/Invalid %s command with %s data, try i.e %s colorname", dataType, data, regNumbersWithColor)
		}
		
		if dataType == slotNumbersWithColor {
			return fmt.Sprintf("Incomplete/Invalid %s command with %s data, try i.e %s slotNum", dataType, data, slotNumbersWithColor)
		}

		if dataType == slotNumberWithReg {
			return fmt.Sprintf("Incomplete/Invalid %s command with %s data, try i.e %s regNum", dataType, data, slotNumberWithReg)
		}
	}

	return ""
}

func processCommand(data []string)  int{
	if checkData("", data) != "" {
		return -1
	}

	dataType := strings.ToLower(data[0])
	switch dataType {
	case createAParkingLot:
		if checkData(dataType, data) != "" {
			log.Printf("%s \n", checkData(dataType, data))
			return -1
		}
		numberOfParkingLots, _ := strconv.Atoi(data[1])
		createParkingLot(numberOfParkingLots)
	case parkVehicle:
		parkingSlots = parkAVehicle(parkingSlots, data[1:], totalNumberOfParkingLots)
	case freeSlot:
		if checkData(dataType, data) != "" {
			log.Printf("%s \n", checkData(dataType, data))
			return -1
		}
		slotNumber, _ := strconv.Atoi(data[1])
		freeAParkingSlot(parkingSlots, totalNumberOfParkingLots, slotNumber)
	case statusOfParkingSlots:
		listAllSlotDetails(parkingSlots)
	case regNumbersWithColor:
		if checkData(dataType, data) != "" {
			log.Printf("%s \n", checkData(dataType, data))
			return -1
		}
		findRegOrSlotByColor(parkingSlots, data[1], true)
	case slotNumbersWithColor:
		if checkData(dataType, data) != "" {
			log.Printf("%s \n", checkData(dataType, data))
			return -1
		}
		findRegOrSlotByColor(parkingSlots, data[1], false)
	case slotNumberWithReg:
		if checkData(dataType, data) != "" {
			log.Printf("%s \n", checkData(dataType, data))
			return -1
		}
		findSlotNumberByRegNum(parkingSlots, data[1])
	default:
		fmt.Println("Not a valid command")
		return -1
	}

	return 1
}
