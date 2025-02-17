package internal

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var procEntries = []string{
	"MemTotal",
	"MemFree",
	"MemAvailable",
	"Buffers",
	"Cached",
	"SwapCached",
	"SwapTotal",
	"SwapFree",
}

var procEntriesMap = map[string]string{}

func Meminfo() map[string]string {
	memInfo, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatalf("Failed to read /proc/meminfo: %v", err)
	}
	defer memInfo.Close()

	scanner := bufio.NewScanner(memInfo)
	for scanner.Scan() {
		line := scanner.Text()

		for _, entry := range procEntries {
			if strings.HasPrefix(line, entry+":") {
				parts := strings.Fields(line)

				if len(parts) >= 2 {
					procEntriesMap[entry] = parts[1]
				}
			}
		}
	}

	return procEntriesMap
}
