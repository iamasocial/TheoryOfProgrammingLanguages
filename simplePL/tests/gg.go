package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i == 10 {
			i = 50
		}
		fmt.Println(i)
	}
}
