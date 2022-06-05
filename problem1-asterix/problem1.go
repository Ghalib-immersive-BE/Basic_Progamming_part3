package main

import "fmt"

func PlayWithAsterix(n int) string {
	// your code her
	var rows int = 5
	var k int

	for i := 1; i <= rows; i++ {
		k = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}
	return
}
func main() {
	fmt.Println(PlayWithAsterix(5))
}
