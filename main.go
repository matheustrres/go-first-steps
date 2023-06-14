package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	monitoringDelay = 5
	logsTextFileName = "tmp/logs.txt"
	websitesTextFileName = "tmp/websites.txt"
)

func main() {
	introduceApplication()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printMonitoringLogs()
		case 0:
			fmt.Println("Current program version", 1.1)
			fmt.Println("Quiting program...")
			os.Exit(0)
		default:
			fmt.Println("No valid instruction specified!")
			os.Exit(-1)
		}
	}
}

func introduceApplication() {
	name := "Matheus"

	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("Hello, %s! Please, choose a command instruction to be executed:", name)
	fmt.Println()
}

func showMenu() {
	fmt.Println()
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Print monitoring logs")
	fmt.Println("0- Quit")
	fmt.Println()
}

func readCommand() int {
	var readedCommand int

	fmt.Scan(&readedCommand)
	fmt.Println("The choosen command was:", readedCommand)
	fmt.Println()

	return readedCommand
}

func startMonitoring() {
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

func printMonitoringLogs() {
	fmt.Println("Printing logs...")
	fmt.Println()

	file, err := ioutil.ReadFile(logsTextFileName)

	if err!= nil {
    fmt.Println("An error has been found while trying printing monitoring logs:", err)
  }

	fmt.Println(string(file))
}

func checkWebsiteStatus(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has been found while trying checking website status:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("The website", site, "is UP! Current status code:", response.StatusCode)
		createMonitoringLog(site, true)
	} else {
		fmt.Println("The website", site, "is DOWN! Current status code:", response.StatusCode)
		createMonitoringLog(site, false)
	}
}

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

func createMonitoringLog(site string, status bool) {
	file, err := os.OpenFile(logsTextFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		fmt.Println("An error has been found while trying opening log file:", err)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	statusFormatted := strconv.FormatBool(status)

	file.WriteString(now + " - " + site + " - online: " + statusFormatted + "\n")

	file.Close()
}