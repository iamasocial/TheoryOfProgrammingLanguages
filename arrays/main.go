package main

import (
	"arrays/opers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("arrays.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Fields(line)
		command[0] = strings.ToLower(command[0])
		if len(command) == 0 {
			continue
		}
		if function, exists := opers.FuncMap[command[0]]; exists {
			function(command, opers.Arrays)
		} else {
			fmt.Println("Invalid command")
			return
		}
	}
	// for {
	// 	// fmt.Print("Enter command: ")
	// 	input := funcs.GetUserInput()
	// 	command := strings.Fields(input)
	// 	if len(command) == 0 {
	// 		fmt.Println("no command entered")
	// 		continue
	// 	}
	// 	command[0] = strings.ToLower(command[0])

	// 	if function, exists := opers.FuncMap[command[0]]; exists {
	// 		function(command, opers.Arrays)
	// 		continue
	// 	}

	// 	fmt.Printf("Unknow command %s\n", command[0])

	// }

}
