package internals

import (
	"fmt"
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

func processCommand(data []string) int {

	dataType := strings.ToLower(data[0])
	switch dataType {
	case createAParkingLot:
		numberOfParkingLots, _ := strconv.Atoi(data[1])
		createParkingLot(numberOfParkingLots)
	case parkVehicle:
		parkingSlots = parkAVehicle(parkingSlots, data[1:], totalNumberOfParkingLots)
	case freeSlot:
		slotNumber, _ := strconv.Atoi(data[1])
		freeAParkingSlot(parkingSlots, totalNumberOfParkingLots, slotNumber)
	case statusOfParkingSlots:
		listAllSlotDetails(parkingSlots)
	case regNumbersWithColor:
		findRegOrSlotByColor(parkingSlots, data[1], true)
	case slotNumbersWithColor:
		findRegOrSlotByColor(parkingSlots, data[1], false)
	case slotNumberWithReg:
		findSlotNumberByRegNum(parkingSlots, data[1])
	default:
		fmt.Println("Not a valid command")
		return -1
	}

	return 1
}
