package main

import "fmt"

/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	pairs := [][]int32{}
	positions := [][]int32{}

	for i := 0; i < len(ar); i++ {
		a := ar[i]

		for j := 0; j < len(ar); j++ {
			b := ar[j]

			if i < j && (a+b)%k == 0 {
				pairs = append(pairs, []int32{a, b})
				positions = append(positions, []int32{int32(i), int32(j)})
			}
		}

	}

	fmt.Println(pairs)
	fmt.Println(positions)

	return int32(len(pairs))

}

func main() {

	numbers := []int32{1, 3, 2, 6, 1, 2}
	k := 3
	n := 6

	result := divisibleSumPairs(int32(n), int32(k), numbers)

	fmt.Println(result)

}
