package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const timesMonitoring = 5
const delaySecondsTimeMonitoring = 5
const nameFile = "sites.txt"

func main() {

	intro()
	for {
		menu()
		command := readCommand()

		switch command {
		case 1:
			monitor()
		case 2:
			readLogs()
		case 0:
			fmt.Println("going out ...")
			os.Exit(0)
		default:
			fmt.Println("ERROR - command invalid!")
			os.Exit(-1)
		}
	}

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

func monitor() {
	fmt.Println("monitoring ...")

	for i := 0; i < timesMonitoring; i++ {
		sites := readFile()

		for _, site := range sites {
			checkSite(site)
		}
		time.Sleep(delaySecondsTimeMonitoring * time.Second)
	}

	fmt.Println()
}

func checkSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("an error ocurred when access site,", site)
	}

	if resp.StatusCode == 200 {
		fmt.Println("site:", site, "loaded successfully!")
		log(site, true)
	} else {
		fmt.Println("an error occurred when access:", site, "statuscode:", resp.StatusCode)
		log(site, false)
	}
}

func fakeSite() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)

	if random%2 == 0 {
		return "https://httpstat.us/500"
	} else {
		return "https://httpstat.us/200"
	}
}

func readFile() []string {
	sites := []string{}

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("an error ocurred when read file,", nameFile, err)
	}

	reader := bufio.NewReader(file)
	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		sites = append(sites, row)

		if err == io.EOF {
			break
		}

	}

	file.Close()

	return append(sites, fakeSite())
}

func log(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("an error ocurred when create log file,", nameFile, err)
	}

	file.WriteString(time.Now().Format("2006/02/01 15:05:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func readLogs() {
	fmt.Println("logs ...")

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("an error ocurred when read log file:", err)
	}

	fmt.Println(string(file))
}
