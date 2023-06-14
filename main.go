package main

import (
	"fmt"
	"os"

	"hello/handlers"
)

func main() {
	introduceApplication()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			handlers.StartMonitoring()
		case 2:
			handlers.PrintMonitoringLogs()
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
