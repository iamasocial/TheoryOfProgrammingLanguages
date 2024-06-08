package main

import (
	"arrays/funcs"
	"arrays/opers"
	"fmt"
	"strings"
)

func main() {

	for {
		fmt.Print("Enter command: ")
		input := funcs.GetUserInput()
		command := strings.Fields(input)
		if len(command) == 0 {
			fmt.Println("no command entered")
			continue
		}
		command[0] = strings.ToLower(command[0])

		if function, exists := opers.FuncMap[command[0]]; exists {
			function(command, opers.Arrays)
			continue
		}

		fmt.Printf("Unknow command %s\n", command[0])

	}

}
