/*
	Complete the function that takes a non-negative integer n as input, and returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).

Examples
n = 0  ==> [1]        # [2^0]
n = 1  ==> [1, 2]     # [2^0, 2^1]
n = 2  ==> [1, 2, 4]  # [2^0, 2^1, 2^2]
*/
package main

import "fmt"

func main() {
	fmt.Println(PowersOfTwo(0))
	fmt.Println(PowersOfTwo(1))
	fmt.Println(PowersOfTwo(2))
}

func PowersOfTwo(n int) []uint64 {

	var result []uint64
	for i := 0; i <= n; i++ {
		result = append(result, uint64(1<<i))
	}
	return result
}
