package main

import "fmt"

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	breakingRecords := [2]int32{0, 0}
	recordMax := scores[0]
	recordMin := scores[0]

	for i := 0; i < len(scores); i++ {
		if scores[i] > recordMax {
			recordMax = scores[i]

			breakingRecords[0]++
		}

		if scores[i] < recordMin {
			recordMin = scores[i]

			breakingRecords[1]++
		}

	}

	return breakingRecords[:]
}

func main() {

	scores := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	result := breakingRecords(scores)

	fmt.Println(result)
}
