package utils

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type MemoryInfo struct {
	Wired        float32
	Active       float32
	Compressed   float32
	Free         float32
	TopProcesses []string
}

// Getting data from vm_stat tool
var getMemoryStatsFromVmStats = func() (map[string]int, error) {
	var memStatMap map[string]int = make(map[string]int)

	outputVmStatCmd, err := exec.Command("vm_stat").Output()

	if err != nil {
		log.Fatal(err)
		return memStatMap, err
	}

	memStatsRows := strings.Split(string(outputVmStatCmd), "\n")

	// Skipping first and last line (header and blank space)
	for _, memStatRow := range memStatsRows[1 : len(memStatsRows)-1] {
		memStatCells := strings.Split(memStatRow, ":")

		memStatkey := strings.TrimSpace(memStatCells[0])
		memStatValue := strings.TrimSpace(strings.Trim(memStatCells[1], "."))

		memStatMap[memStatkey], _ = strconv.Atoi(memStatValue)
	}

	return memStatMap, nil
}

var convertMemoryPagesToGigabyte = func(pages int) float32 {

	// Page size of 4096 bytes
	const PageSize = 4096
	const MemUnit = 1024
	const baseDecimal = 100

	bytesPerPage := pages
	bytesPerPage = bytesPerPage * PageSize

	memoryOnBytes := float64(bytesPerPage) / MemUnit / MemUnit / MemUnit
	memoryOnGBSize := float32(int(memoryOnBytes*baseDecimal)) / baseDecimal

	return memoryOnGBSize
}

// TODO: Change later to a common file and add filter as parameter
var getTopFiveProcessesByMemory = func() []string {

	// $1 = USER, $2 = PID, $3 = %CPU, $4 = %MEM, $11 = COMMAND
	outputTopCmd, err := exec.Command("bash", "-c",
		"ps aux | awk '{print $3, $4, $11}' | sort -rk 2 | head -n 5").Output()

	if err != nil {
		log.Fatal(err)
		return []string{}
	}

	processesRows := strings.Split(string(outputTopCmd), "\n")

	// Second index from print (awk command with $11)
	const ProcessNameIndex = 2
	var topProcesses = make([]string, 5)

	// Skipping last line (blank space)
	for index, processRow := range processesRows[:len(processesRows)-1] {
		processCells := strings.Split(processRow, " ")

		processPaths := strings.Split(processCells[ProcessNameIndex], "/")

		topProcesses[index] = processPaths[len(processPaths)-1]
	}

	return topProcesses
}

func GetMemoryInfo() MemoryInfo {

	const WiredMemKey = "Pages wired down"
	const ActiveMemKey = "Pages active"
	const CompressedMemKey = "Pages occupied by compressor"
	const InactiveMemKey = "Pages inactive"
	const SpeculativeMemKey = "Pages speculative"
	const FreeMemKey = "Pages free"

	memoryStats, err := getMemoryStatsFromVmStats()

	if err != nil {
		log.Fatal(err)
		return MemoryInfo{}
	}

	wiredMemory := convertMemoryPagesToGigabyte(memoryStats[WiredMemKey])
	activeMemory := convertMemoryPagesToGigabyte(memoryStats[ActiveMemKey])
	compressedMemory := convertMemoryPagesToGigabyte(memoryStats[CompressedMemKey])
	freeMemory := convertMemoryPagesToGigabyte(memoryStats[InactiveMemKey]) +
		convertMemoryPagesToGigabyte(memoryStats[SpeculativeMemKey]) +
		convertMemoryPagesToGigabyte(memoryStats[FreeMemKey])

	topProcesses := getTopFiveProcessesByMemory()

	return MemoryInfo{
		Wired:        wiredMemory,
		Active:       activeMemory,
		Compressed:   compressedMemory,
		Free:         freeMemory,
		TopProcesses: topProcesses,
	}
}
