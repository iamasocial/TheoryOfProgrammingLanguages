package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	level := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if char != "#" && char != "{" && char != "}" && level == 0 {
				fmt.Print(char)
			}

			if char == "}" && level <= 0 {
				fmt.Println("\n!!!!!an error occurred while reading the file!!!!!")
				os.Exit(1)
			}

			if char == "#" {
				i = len(line)
				continue
			}

			if char == "{" {
				level++
				continue
			}

			if char == "}" {
				level--
				continue
			}
		}
		fmt.Print("\n")
	}
}
