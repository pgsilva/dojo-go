package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	// ordenando o array asc
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	sets := [][]int32{}
	bird := int32(0)

	//agrupando os passaros
	for i := 0; i < len(arr); i++ {
		if arr[i] != bird {
			bird = arr[i]
			types := []int32{}

			for j := 0; j < len(arr); j++ {
				if bird == arr[j] {
					types = append(types, arr[j])
				}
			}

			sets = append(sets, types)
		}
	}

	ids := []int32{}
	maxBirds := 0
	//maiores conjuntos de passaros
	for _, val := range sets {
		if len(val) > maxBirds {
			maxBirds = len(val)
		}
	}

	for idx, val := range sets {
		if len(val) == maxBirds {
			ids = append(ids, int32(idx))
		}
	}

	//o menor id de passaro dentro dos maiores conjuntos
	minIdx := int32(6)
	for _, val := range ids {
		if val < int32(minIdx) {
			minIdx = val
		}
	}

	return sets[minIdx][0]
}

func main() {

	birds := []int32{1, 4, 4, 4, 5, 3}
	result := migratoryBirds(birds)

	fmt.Println(result)
}
