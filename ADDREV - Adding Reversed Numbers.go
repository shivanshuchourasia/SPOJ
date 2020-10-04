package main

import (
	"fmt"
	"strconv"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		fmt.Println(revInt(revInt(a) + revInt(b)))
		t--
	}
}

func revInt(x int) int {
	X := strconv.Itoa(x)
	isLeadingZero := true
	Y := ""
	for i := len(X) - 1; i >= 0; i-- {
		if (!isLeadingZero && X[i] == '0') || X[i] != '0' {
			isLeadingZero = false
			Y += string(X[i])
		}
	}
	if ans, err := strconv.Atoi(Y); err == nil {
		return ans
	}
	return x
}
