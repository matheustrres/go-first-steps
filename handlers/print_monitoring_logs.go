package handlers

import (
	"fmt"
	"io/ioutil"
)

func PrintMonitoringLogs() {
	fmt.Println("Printing logs...")
	fmt.Println()

	file, err := ioutil.ReadFile(logsTextFileName)

	if err != nil {
		fmt.Println("An error has been found while trying printing monitoring logs:", err)
	}

	fmt.Println(string(file))
}
