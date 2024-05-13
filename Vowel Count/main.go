/*
	Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.
*/

package main

import "fmt"

func main() {
	fmt.Println(GetCount("Hello World"))
}

func GetCount(str string) (count int) {
	for _, v := range str {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			count++
		}
	}
	return count
}
