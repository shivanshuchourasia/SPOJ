package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		var n int
		fmt.Scanf("%d", &n)
		var x, ans, i int = n, 0, 5
		for x != 0 {
			x = n / i
			ans += x
			i *= 5
		}
		fmt.Println(ans)
		t--
	}
}
