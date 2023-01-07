package main

import "fmt"

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	countCandles := 0
	maxHeightCandles := 0

	for i := 0; i < len(candles); i++ {
		if candles[i] > int32(maxHeightCandles) {
			maxHeightCandles = int(candles[i])
		}
	}

	for i := 0; i < len(candles); i++ {
		if candles[i] == int32(maxHeightCandles) {
			countCandles++
		}
	}

	fmt.Println(countCandles)

	return int32(countCandles)
}

func main() {
	candles := []int32{3, 2, 1, 3}
	birthdayCakeCandles(candles[:])
}
