package main

import "fmt"

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 * s: integer, starting point of Sam's house location.
 * t: integer, ending location of Sam's house location.
 * a: integer, location of the Apple tree.
 * b: integer, location of the Orange tree.
 * apples: integer array, distances at which each apple falls from the tree.
 * oranges: integer array, distances at which each orange falls from the tree.
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	distancesApples := []int32{}
	distancesOranges := []int32{}

	totalApples := 0
	totalOranges := 0

	for i := 0; i < len(apples); i++ {
		distancesApples = append(distancesApples, a+apples[i])
	}

	for i := 0; i < len(oranges); i++ {
		distancesOranges = append(distancesOranges, b+oranges[i])
	}

	for i := 0; i < len(distancesApples); i++ {
		if s <= distancesApples[i] && distancesApples[i] <= t {
			totalApples++
		}
	}

	for i := 0; i < len(distancesOranges); i++ {
		if s <= distancesOranges[i] && distancesOranges[i] <= t {
			totalOranges++
		}
	}

	fmt.Println(totalApples)
	fmt.Println(totalOranges)

}

func main() {

	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

	countApplesAndOranges(7, 11, 5, 15, apples, oranges)

}
