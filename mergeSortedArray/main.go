package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) > m {
		nums1 = nums1[:m]
	}
	if len(nums2) > n {
		nums2 = nums2[:n]
	}
	for _, v := range nums2 {
		nums1 = append(nums1, v)
	}
	sort.Ints(nums1)
}

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}
