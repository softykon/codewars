package main

import "fmt"

func main() {
	fmt.Println(IsTriangle(1, 2, 2))
	fmt.Println(IsTriangle(7, 2, 2))
	fmt.Println(IsTriangle(1, 2, 3))
}
func IsTriangle(a, b, c int) bool {

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return false
	}
	return true
}
