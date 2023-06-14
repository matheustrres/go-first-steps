package handlers

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	logsTextFileName = "tmp/logs.txt"
)

func CreateMonitoringLog(site string, status bool) {
	file, err := os.OpenFile(logsTextFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		fmt.Println("An error has been found while trying opening log file:", err)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	statusFormatted := strconv.FormatBool(status)

	file.WriteString(now + " - " + site + " - online: " + statusFormatted + "\n")

	file.Close()
}
