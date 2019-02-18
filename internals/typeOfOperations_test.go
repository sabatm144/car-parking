package internals

import (
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
		name       string
		slotMap    map[int]vehicleInfo
		color      string
		regOrColor bool
		want       int
	}{
		{"Valid- Type 1", map[int]vehicleInfo{1: vehicleInfo{Number: "reg1", Color: "W"}}, "W", true, 1},
		{"Valid- Type 2", map[int]vehicleInfo{2: vehicleInfo{Number: "reg2", Color: "B"}}, "B", false, 1},
		{"InValid- Type 1", map[int]vehicleInfo{3: vehicleInfo{Number: "reg3", Color: ""}}, "W", true, -1},
		{"Valid- Type 2", map[int]vehicleInfo{4: vehicleInfo{Number: "reg4", Color: ""}}, "W", false, -1},
		{"Valid- Type 3", nil, "W", false, -1},
	}

	for _, tt := range tests {
		result := findRegOrSlotByColor(tt.slotMap, tt.color, tt.regOrColor)
		if result != tt.want {
			t.Errorf("Vehicle with regNum %s not found,  want: %d.", tt.name, tt.want)
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
	}
	for _, tt := range tests {
		parkAVehicle(tt.parkingSlots, tt.data, tt.totalSlots)
	}
}

func Test_checkData(t *testing.T) {
	tests := []struct {
		operationType string
		data          []string
		want          int
	}{
		{createAParkingLot, []string{createAParkingLot}, -1},
		{createAParkingLot, []string{createAParkingLot, "6"}, 1},
		{freeSlot, []string{freeSlot, "6"}, 1},
		{freeSlot, []string{freeSlot}, -1},
		{regNumbersWithColor, []string{regNumbersWithColor}, -1},
		{regNumbersWithColor, []string{regNumbersWithColor, "W"}, 1},
		{slotNumbersWithColor, []string{slotNumbersWithColor}, -1},
		{slotNumbersWithColor, []string{slotNumbersWithColor, "W"}, 1},
		{slotNumberWithReg, []string{slotNumberWithReg}, -1},
		{slotNumberWithReg, []string{slotNumberWithReg, "reg1"}, 1},
	}
	for _, tt := range tests {
		result := checkData(tt.data)
		if result != tt.want {
			t.Errorf("Invalid command data for %s with %d want %d", tt.operationType, result, tt.want)
		}
	}
}

func Test_processCommand(t *testing.T) {
	tests := []struct {
		data []string
		want int
	}{
		{[]string{"not a valid command"}, -1},
		{[]string{createAParkingLot, "6"}, 1},
		{[]string{parkVehicle, "KA-01-HH-1234", "W"}, 1},
		{[]string{freeSlot, "4"}, 1},
		{[]string{statusOfParkingSlots}, 1},
		{[]string{regNumbersWithColor, "W"}, 1},
		{[]string{slotNumberWithReg, "KA-01-HH-1234"}, 1},
	}
	for _, tt := range tests {
		result := processCommand(tt.data)
		if result != tt.want {
			t.Errorf("Invalid result %d want %d", result, tt.want)
		}
	}

}
