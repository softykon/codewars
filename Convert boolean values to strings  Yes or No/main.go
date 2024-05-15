/* Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false. */

package main

import "fmt"

func main() {

	fmt.PrintInt(BoolToWord(true))
}

func BoolToWord(word bool) string {

	if word == true {
		return "Yes"
	} else {
		return "No"
	}
}
