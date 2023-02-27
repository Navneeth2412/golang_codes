package main

import "fmt"

func main() {
	var maps = map[string]int{}
	var n int
	fmt.Println("Enter the number of values to be added in map.")
	fmt.Scanf("%d", &n)
	fmt.Println("Enter the keys and values: ")
	for i := 0; i < n; i++ {
		var num int
		var text string
		fmt.Scanf("%s %d", &text, &num)
		maps[text] = num
	}
	fmt.Println(maps)

}
