package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	limit := 38

	gradesRounded := []int32{}

	for i := 0; i < len(grades); i++ {
		if grades[i] < int32(limit) {
			gradesRounded = append(gradesRounded, grades[i])
			continue
		}

		if grades[i]%5 == 3 {
			gradesRounded = append(gradesRounded, grades[i]+2)
		} else if grades[i]%5 == 4 {
			gradesRounded = append(gradesRounded, grades[i]+1)
		} else {
			gradesRounded = append(gradesRounded, grades[i])
		}

	}

	fmt.Println(gradesRounded)
	return gradesRounded
}

func main() {

	grades := []int32{73, 67, 38, 33}
	gradingStudents(grades)

}
