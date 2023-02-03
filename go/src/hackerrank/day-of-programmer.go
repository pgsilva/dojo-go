package main

import (
	"fmt"
	"strconv"
)

func dayOfProgrammer(year int32) string {
	// Write your code here

	if year == 1918 {
		return "26.09.1918"
	} else if year <= 1917 && year%4 == 0 || year%400 == 0 || year%4 == 0 && year%100 != 0 {
		return "12.09." + strconv.Itoa(int(year))
	} else {
		return "13.09." + strconv.Itoa(int(year))
	}

}

func main() {

	result := dayOfProgrammer(2016)

	fmt.Println(result)
}
