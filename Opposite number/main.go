/* Very simple, given a number (integer / decimal / both depending on the language), find its opposite (additive inverse).

Examples:

1: -1
14: -14
-34: 34 */

package main

import "fmt"

func main() {
	fmt.Println(Opposite(1))
	fmt.Println(Opposite(14))
	fmt.Println(Opposite(-34))
}

func Opposite(value int) int {

	return -value

}
