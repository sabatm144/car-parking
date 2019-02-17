package internals

import (
	"fmt"
	"log"
	"strings"
)

var (
	assingnedParkingLots int
)

func createParkingLot(parkingLot int) int {

	if totalNumberOfParkingLots > 0 {
		log.Println("Parking lot can be created only once!")
		return -1
	}

	if parkingLot > 0 {
		totalNumberOfParkingLots = parkingLot
		fmt.Printf("Created a parking lot with %d slots \n", parkingLot)
		return parkingLot
	}

	return -1
}

func parkAVehicle(parkingSlots map[int]vehicleInfo, data []string, parkingLots int) map[int]vehicleInfo {

	if len(data) < 2 || len(data) > 2 {
		println("Incomplete/Invalid park command, try i.e park regNum solNum")
		return nil
	}

	regNum := data[0]
	color := data[1]
	for slotNum, vehicle := range parkingSlots {
		if !vehicle.Alloted {
			parkingSlots[slotNum] = vehicleInfo{Number: regNum, Color: color, Alloted: true}
			fmt.Printf("Alloted Slot No. %d \n", slotNum)
			return parkingSlots
		}
	}

	if assingnedParkingLots >= parkingLots {
		fmt.Println("Sorry, Parking lot is full")
		return parkingSlots
	}
	assingnedParkingLots = assingnedParkingLots + 1
	// log.Println("parking Lot :", assingnedParkingLots)

	if parkingSlots == nil {
		parkingSlots = make(map[int]vehicleInfo)
	}
	parkingSlots[assingnedParkingLots] = vehicleInfo{Number: regNum, Color: color, Alloted: true}
	fmt.Printf("Alloted Slot No.. %d \n", assingnedParkingLots)
	return parkingSlots
}

func freeAParkingSlot(parkingSlots map[int]vehicleInfo, parkingLots, slotNumber int) int {
	if slotNumber > 0 && parkingLots > 0 && slotNumber <= parkingLots {
		if parkingSlots == nil {
			parkingSlots = make(map[int]vehicleInfo)
		}
		parkingSlots[slotNumber] = vehicleInfo{}
		fmt.Printf("Slot No. %d is free \n", slotNumber)
		return slotNumber
	}
	fmt.Printf("Invalid slot No.: %d \n", slotNumber)
	return -1
}

func listAllSlotDetails(parkingSlots map[int]vehicleInfo) int {
	if parkingSlots == nil {
		fmt.Println("No cars parked !")
		return -1
	}
	fmt.Println("SlotNum:  RegNum:  Color: ")
	for slotNum, vehicle := range parkingSlots {
		if vehicle.Number != "" {
			fmt.Printf("%d     %s    %s \n", slotNum, vehicle.Number, vehicle.Color)
		}
	}

	fmt.Println("")
	return 1
}

func findRegOrSlotByColor(parkingSlots map[int]vehicleInfo, color string, regOrSlot bool) int {
	found := false
	for slotNum, vehicle := range parkingSlots {
		if strings.EqualFold(vehicle.Color, color) && regOrSlot {
			found = true
			fmt.Printf("%s, ", vehicle.Number)
		}
		if strings.EqualFold(vehicle.Color, color) && !regOrSlot {
			found = true
			fmt.Printf("%d, ", slotNum)
		}
	}
	if !found && len(parkingSlots) > 0 {
		fmt.Println("Not found")
		return -1
	}
	fmt.Println("")
	return 1
}

func findSlotNumberByRegNum(parkingSlots map[int]vehicleInfo, regNo string) int {
	found := false
	for slotNum, vehicle := range parkingSlots {
		if strings.EqualFold(vehicle.Number, regNo) {
			found = true
			fmt.Printf("%d, ", slotNum)
			break
		}
	}
	if !found {
		fmt.Println("Not found")
		return -1
	}
	fmt.Println("")
	return 1
}
