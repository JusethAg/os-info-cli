package utils

import (
	"reflect"
	"testing"
)

var memoryInfo, _ = getMemoryStatsFromVmStats()
var wiredMemoryInfo = memoryInfo["Pages wired down"]
var activeMemoryInfo = memoryInfo["Pages active"]
var compressedMemoryInfo = memoryInfo["Pages occupied by compressor"]
var inactiveMemoryInfo = memoryInfo["Pages inactive"]
var speculativeMemoryInfo = memoryInfo["Pages speculative"]
var freeMemoryInfo = memoryInfo["Pages free"]

// TODO: Update tests refactoring implementation with interfaces and structs
func TestGetMemoryInfo(t *testing.T) {
	tests := []struct {
		name string
		want MemoryInfo
	}{
		{
			"Get memory details with top 5 processes",
			MemoryInfo{
				Wired:      convertMemoryPagesToGigabyte(wiredMemoryInfo),
				Active:     convertMemoryPagesToGigabyte(activeMemoryInfo),
				Compressed: convertMemoryPagesToGigabyte(compressedMemoryInfo),
				Free: convertMemoryPagesToGigabyte(inactiveMemoryInfo) +
					convertMemoryPagesToGigabyte(speculativeMemoryInfo) +
					convertMemoryPagesToGigabyte(freeMemoryInfo),
				TopProcesses: getTopFiveProcessesByMemory(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMemoryInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMemoryInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
