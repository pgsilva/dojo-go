package main

import (
	"fmt"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	if x1 == x2 {
		return "YES"
	}

	var diff int32 = 0
	if v1 > v2 {
		diff = v1 - v2
	} else {
		diff = v2 - v1
	}

	if diff == 0 {
		return "NO"
	}

	xdiff := x1 - x2
	vdiff := v2 - v1

	if xdiff < 0 && vdiff < 0 || xdiff > 0 && vdiff > 0 {
		mod := xdiff % vdiff
		mod2 := vdiff % xdiff

		if mod == 0 || mod2 == 0 {
			return "YES"
		}
	}

	return "NO"
}

func main() {

	res := kangaroo(0, 3, 4, 2)
	fmt.Println(res)
}
