package main

import "fmt"

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	count := int32(0)

	for i := 0; i < len(s)-int(m-1); i++ {
		sum := int32(0)

		for j, k := i, 0; int32(k) < m; j, k = j+1, k+1 {
			sum += s[j]
		}

		if sum == d {
			count++
		}
	}

	return count
}

func main() {

	squares := []int32{4}
	result := birthday(squares, 4, 1)

	fmt.Println(result)
}
