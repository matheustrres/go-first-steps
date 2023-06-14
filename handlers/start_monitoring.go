package handlers

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	monitoringDelay      = 5
	websitesTextFileName = "tmp/websites.txt"
)

func readWebsitesFromTextFile() []string {
	var sites []string

	file, err := os.Open(websitesTextFileName)

	if err != nil {
		fmt.Println("An error has been found while trying reading text file:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func checkWebsiteStatus(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has been found while trying checking website status:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("The website", site, "is UP! Current status code:", response.StatusCode)
		CreateMonitoringLog(site, true)
	} else {
		fmt.Println("The website", site, "is DOWN! Current status code:", response.StatusCode)
		CreateMonitoringLog(site, false)
	}
}

func StartMonitoring() {
	fmt.Println("Starting monitoring...")

	sites := readWebsitesFromTextFile()

	numberOfMonitorings := len(sites)

	for i := 0; i < numberOfMonitorings; i++ {
		for _, site := range sites {
			checkWebsiteStatus(site)
		}

		time.Sleep(monitoringDelay * time.Second)
		fmt.Println()
	}

	fmt.Println("----------------------------------------------------------------------")
}
