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
