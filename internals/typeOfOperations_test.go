package internals

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func Test_createParkingLot(t *testing.T) {
	tests := []struct {
		lot  int
		want int
	}{
		{4, 4},
		{0, -1},
		{-1, -1},
	}
	for _, tt := range tests {
		result := createParkingLot(tt.lot)
		if result != tt.want {
			t.Errorf("Invalid Parking Lot  %d,  want: %d.", result, tt.want)
		}
	}
}

func Test_freeAParkingSlot(t *testing.T) {
	tests := []struct {
		pLot       int
		parkingMap map[int]vehicleInfo
		slot       int
		want       int
	}{
		{4, map[int]vehicleInfo{4: vehicleInfo{Number: "Reg1", Color: "White"}}, 5, -1},
		{4, map[int]vehicleInfo{4: vehicleInfo{Number: "Reg2", Color: "Blue"}}, 4, 4},
		{4, nil, 4, 4},
	}
	for _, tt := range tests {
		result := freeAParkingSlot(tt.parkingMap, tt.pLot, tt.slot)
		if result != tt.want {
			t.Errorf("Invalid Parking Lot with slot %d,  want: %d.", result, tt.want)
		}
	}
}
func Test_findSlotNumberByRegNum(t *testing.T) {
	tests := []struct {
		slot  int
		vInfo vehicleInfo
		regNo string
		want  int
	}{
		{1, vehicleInfo{Number: "reg1", Color: "W"}, "reg1", 1},
		{2, vehicleInfo{Number: "reg2", Color: "W"}, "", -1},
		{2, vehicleInfo{}, "reg2", -1},
	}
	for _, tt := range tests {
		slotMap := make(map[int]vehicleInfo)
		slotMap[tt.slot] = tt.vInfo
		result := findSlotNumberByRegNum(slotMap, tt.regNo)
		if result != tt.want {
			t.Errorf("Vehicle with regNum %s not found,  want: %d.", tt.vInfo.Number, tt.want)
		}
	}
}
func Test_findRegOrSlotByColor(t *testing.T) {
	tests := []struct {
		slot       int
		vInfo      vehicleInfo
		regOrColor bool
		want       int
	}{
		{1, vehicleInfo{Number: "reg1", Color: "W"}, true, 1},
		{2, vehicleInfo{Number: "reg2", Color: "B"}, false, 1},
		{3, vehicleInfo{Number: "reg3", Color: ""}, true, -1},
		{4, vehicleInfo{Number: "reg4", Color: ""}, false, -1},
		{5, vehicleInfo{}, false, -1},
	}
	for _, tt := range tests {
		slotMap := make(map[int]vehicleInfo)
		slotMap[tt.slot] = tt.vInfo
		result := findRegOrSlotByColor(slotMap, tt.vInfo.Color, tt.regOrColor)
		if result != tt.want {
			if tt.regOrColor {
				t.Errorf("Vehicle with regNum %s not found,  want: %d.", tt.vInfo.Number, tt.want)
			}
			if !tt.regOrColor {
				t.Errorf("Vehicle with slot %d not found,  want: %d.", tt.slot, tt.want)
			}
		}
	}
}

func Test_listAllSlotDetails(t *testing.T) {
	tests := []struct {
		slotMap map[int]vehicleInfo
		want    int
	}{
		{nil, -1},
		{map[int]vehicleInfo{2: vehicleInfo{Number: "reg2", Color: "B"}}, 1},
	}
	for _, tt := range tests {
		result := listAllSlotDetails(tt.slotMap)
		if result != tt.want {
			t.Errorf("Vehicle with slots not found,  want: %d.", tt.want)
		}
	}
}

func Test_parkAVehicle(t *testing.T) {
	tests := []struct {
		totalSlots   int
		parkingSlots map[int]vehicleInfo
		data         []string
		want         map[int]vehicleInfo
	}{
		{1, map[int]vehicleInfo{2: vehicleInfo{Number: "", Color: "", Alloted: false}}, []string{"reg1", "W"}, map[int]vehicleInfo{2: vehicleInfo{Number: "reg1", Color: "W", Alloted: true}}},
		{1, map[int]vehicleInfo{2: vehicleInfo{Number: "", Color: "", Alloted: false}}, []string{}, nil},
		{1, nil, []string{"reg1", "W"}, map[int]vehicleInfo{1: vehicleInfo{Number: "reg1", Color: "W", Alloted: true}}},
	}
	for _, tt := range tests {
		parkAVehicle(tt.parkingSlots, tt.data, tt.totalSlots)
	}
}

func Test_checkData(t *testing.T) {
	tests := []struct {
		operationType string
		data          []string
		want          string
	}{
		{"", []string{}, "Command not found"},
		{createAParkingLot, []string{createAParkingLot}, fmt.Sprintf("Incomplete/Invalid %s command with %s data, try i.e %s num", createAParkingLot, []string{createAParkingLot}, createAParkingLot)},
		{createAParkingLot, []string{createAParkingLot, "6"}, ""},
		{freeSlot, []string{freeSlot, "6"}, ""},
		{freeSlot, []string{freeSlot}, fmt.Sprintf("Incomplete/Invalid %s command, try i.e %s solNum", []string{freeSlot}, freeSlot)},
		{regNumbersWithColor, []string{regNumbersWithColor}, fmt.Sprintf("Incomplete/Invalid %s command with %s data, try i.e %s colorname", regNumbersWithColor, []string{regNumbersWithColor}, regNumbersWithColor)},
		{regNumbersWithColor, []string{regNumbersWithColor, "W"}, ""},
		{slotNumbersWithColor, []string{slotNumbersWithColor}, fmt.Sprintf("Incomplete/Invalid %s command with %s data, try i.e %s slotNum", slotNumbersWithColor, []string{slotNumbersWithColor}, slotNumbersWithColor)},
		{slotNumberWithReg, []string{slotNumberWithReg}, fmt.Sprintf("Incomplete/Invalid %s command with %s data, try i.e %s regNum", slotNumberWithReg, []string{slotNumberWithReg}, slotNumberWithReg)},
	}
	for _, tt := range tests {
		result := checkData(tt.operationType, tt.data)
		if !strings.EqualFold(result, tt.want) {
			log.Println(result)
			log.Println(tt.want)

			t.Errorf("Invalid command data  with %s want %s", result, tt.want)
		}
	}
}

func Test_processCommand(t *testing.T) {
	tests := []struct {
		data []string
		want int
	}{
		{[]string{createAParkingLot, "6"}, 1},
		{[]string{createAParkingLot}, -1},
		{[]string{parkVehicle, "KA-01-HH-1234", "W"}, 1},
		{[]string{freeSlot, "4"}, 1},
		{[]string{freeSlot}, -1},
		{[]string{statusOfParkingSlots}, 1},
		{[]string{regNumbersWithColor}, -1},
		{[]string{regNumbersWithColor, "W"}, 1},
		{[]string{slotNumbersWithColor}, -1},
		{[]string{slotNumberWithReg, "KA-01-HH-1234"}, 1},

	}
	for _, tt := range tests {
		result := processCommand(tt.data)
		if result != tt.want {
			t.Errorf("Invalid result %d want %d", result, tt.want)
		}
	}
	
}
