package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here
	s = strings.ToLower(s)
	pm := strings.HasSuffix(s, "pm")
	am := strings.HasSuffix(s, "am")
	t := s[:len(s)-2]

	timeArr := strings.Split(t, ":")

	hh, mm, ss := timeArr[0], timeArr[1], timeArr[2]
	hhInt, err := strconv.Atoi(hh)

	if err != nil {
		panic(err.Error())
	}

	if am && hhInt == 12 {
		hhInt = 0
	}

	if pm && hhInt != 12 {
		hhInt += 12
	}

	hh = strconv.Itoa(hhInt)

	result := fmt.Sprintf("%02s:%02s:%02s", hh, mm, ss)

	fmt.Println(result)

	return result
}

func main() {
	timeConversion("07:05:45PM")
}
