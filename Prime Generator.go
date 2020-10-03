package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	for t > 0 {
		var a, b int
		fmt.Scanf("%v %v", &a, &b)
		for i := a; i <= b; i++ {
			if isPrime(i) {
				fmt.Printf("%v\n", i)
			}
		}
		// fmt.Printf("\n")
		t--
	}
}

func isPrime(x int) bool {
	if x <= 1 {
		return false
	}
	if x <= 3 {
		return true
	}
	if x%2 == 0 || x%3 == 0 {
		return false
	}
	for i := 5; i*i <= x; i += 6 {
		if x%i == 0 || x%(i+2) == 0 {
			return false
		}
	}
	return true
}
