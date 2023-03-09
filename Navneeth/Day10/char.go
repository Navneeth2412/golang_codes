package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Enter a word")
	var word string
	fmt.Scanf("%s", &word)
	go occ(word)
	time.Sleep(time.Millisecond * 100)
}

func occ(w string) {
	count := 0
	res := []string{}
	for _, v := range w {
		res = append(res, string(v))
	}
	for i := 0; i < len(w); i++ {
		count = 0
		for j := i; j < len(w); j++ {
			if res[j] == res[i] {
				count += 1
			}
		}
		fmt.Println(res[i], "=", count)
	}
	// fmt.Println(count)
}
