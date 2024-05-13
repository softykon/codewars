/* This code does not execute properly. Try to figure out why. */

package main

import "fmt"

func main() {
	fmt.Println(Multiply(2, 3))
}

func Multiply(a, b int) int {
	return a * b
}
