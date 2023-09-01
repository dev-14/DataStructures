package main

import "fmt"

func removeElement(nums []int, val int) int {
	index := 0
	for _, rem := range nums {
		if rem != val {
			nums[index] = rem
			index += 1
		}

	}

	return index
}

func main() {
	numList := []int{3, 2, 2, 3, 1}
	val := 2
	i := removeElement(numList, val)

	fmt.Println(i)
}
