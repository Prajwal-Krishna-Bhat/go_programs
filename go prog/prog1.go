package main

import "fmt"

func main() {
	var prev, current, sum, i int
	var a [10]int
	prev = 0
	current = 1
	sum = 0
	i = 0
	for prev+current < 1000 {
		next := prev + current
		if next%2 == 0 {
			a[i] = next
			i += 1
			sum += next
		}
		prev = current
		current = next
	}
	fmt.Println("The even terms are:", a)
	fmt.Println("the sum of the even valued terms is", sum)
}
