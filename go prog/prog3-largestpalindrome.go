package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
func largestPalindromeProduct() (int, int, int) {
	lp := 0
	var mul1, mul2 int
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			if i*j < lp {
				break
			}
			if i*j > lp && isPalindrome(i*j) {
				lp = i * j
				mul1 = i
				mul2 = j
			}
		}
	}
	return lp, mul1, mul2
}
func main() {
	result, mul1, mul2 := largestPalindromeProduct()
	fmt.Printf("The largest palindrome product is: %d\n", result)
	fmt.Printf("The multiplicants are: %d and %d\n", mul1, mul2)
}
