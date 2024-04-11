package main

import (
	"arrays/funcs"
	"arrays/opers"
	"fmt"
	"os"
	"strings"
)

func main() {
	arrays := make(map[string][]int)

	for {
		fmt.Print("Enter command: ")
		input := funcs.GetUserInput()
		command := strings.Fields(input)
		command[0] = strings.ToLower(command[0])

		switch command[0] {
		case "load":
			opers.LoadArray(command, arrays)
		case "save":
			opers.SaveArray(command, arrays)
		case "rand":
			opers.Random(command, arrays)
		case "concat":
			opers.Concat(command, arrays)
		case "free":
			opers.Free(command, arrays)
		case "remove":
			opers.Remove(command, arrays)
		case "copy":
			opers.Copy(command, arrays)
		case "sort":
			opers.Sort(command, arrays)
		case "shuffle":
			opers.Shuffle(command, arrays)
		case "stats":
			opers.Stats(command, arrays)
		case "print":
			opers.Print(command, arrays)
		case "exit":
			fmt.Println("Bye :)")
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Try again")
		}
	}

}
