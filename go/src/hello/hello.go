package main

import (
	"fmt"
	"os"
)

func main() {

	intro()
	menu()

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("monitoring ...")
	case 2:
		fmt.Println("logs ...")
	case 0:
		fmt.Println("going out ...")
	default:
		fmt.Println("ERROR - command invalid!")
	}

	os.Exit(0)
}

func intro() {
	//another form declare variable
	var name string = "Mr Robot"

	version := 0.1

	fmt.Println(name)
	fmt.Println("version:", version)

}

func menu() {
	fmt.Println("#######")
	fmt.Println()
	fmt.Println("choose any option to start ...")
	fmt.Println("1 - start monitoring")
	fmt.Println("2 - show logs")
	fmt.Println("0 - exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("option selected:", command)

	return command

}
