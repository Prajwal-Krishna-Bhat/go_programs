package main

import "fmt"

func findSpecialTriplet(n int) (int, int, int) {
	for a := 1; a <= n/3; a++ {
		for b := a; b <= (n-a)/2; b++ {
			c := n - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return -1, -1, -1
}
func main() {
	var n int
	fmt.Printf("enter the value of sum\n")
	fmt.Scan(&n)
	a, b, c := findSpecialTriplet(n)
	if a != -1 && b != -1 && c != -1 {
		fmt.Printf("The pythogoream triplet for sum %d is (%d %d %d)\n", n, a, b, c)
	} else {
		fmt.Printf("no special triplet was found for sum %d\n", n)
	}
}
