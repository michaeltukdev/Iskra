package internal

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Uptime() map[string]string {
	cpuInfo, err := os.Open("/host/proc/uptime")
	if err != nil {
		log.Fatalf("Failed to read /host/proc/cpuinfo: %v", err)
	}
	defer cpuInfo.Close()

	scanner := bufio.NewScanner(cpuInfo)

	var uptime string
	var idle string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)

		uptime = parts[0]
		idle = parts[1]
	}

	return map[string]string{
		"Uptime": uptime,
		"Idle":   idle,
	}
}
