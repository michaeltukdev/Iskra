package internal

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var procEntries = []string{
	"model name",
	"cpu MHz",
	"cache size",
	"cpu cores",
	"siblings",
	"physical id",
}

var keyRenames = map[string]string{
	"model name":  "Model Name",
	"cpu MHz":     "Clock Speed MHz",
	"cache size":  "Cache Size",
	"cpu cores":   "Core Count",
	"siblings":    "Thread Count",
	"physical id": "Physical ID",
}

func CPUInfo() map[string]map[string]string {
	processors := make(map[string]map[string]string)

	cpuInfo, err := os.Open("/host/proc/cpuinfo")
	if err != nil {
		log.Fatalf("Failed to read /host/proc/cpuinfo: %v", err)
	}
	defer cpuInfo.Close()

	scanner := bufio.NewScanner(cpuInfo)

	var currentPhysicalID string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		for _, entry := range procEntries {
			if strings.HasPrefix(line, entry) {

				parts := strings.SplitN(line, ":", 2)
				if len(parts) != 2 {
					continue
				}
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				customKey, found := keyRenames[key]
				if !found {
					customKey = key
				}

				if key == "physical id" {
					currentPhysicalID = value

					if _, ok := processors[currentPhysicalID]; !ok {
						processors[currentPhysicalID] = make(map[string]string)
					}
				}

				if currentPhysicalID != "" {
					if _, ok := processors[currentPhysicalID]; !ok {
						processors[currentPhysicalID] = make(map[string]string)
					}
					processors[currentPhysicalID][customKey] = value
				}
			}
		}
	}

	return processors
}
