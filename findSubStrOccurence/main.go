package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

func main() {
	str := "hello world"
	substr := "hell"
	i := strStr(str, substr)

	fmt.Println(i)
}
