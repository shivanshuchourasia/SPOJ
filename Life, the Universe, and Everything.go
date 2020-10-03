package main

import "fmt"

func main() {
	// your code goes here
	for {
		var a int
		fmt.Scanln(&a)
		if a == 42 {
			break
		}
		fmt.Println(a)
	}
}
