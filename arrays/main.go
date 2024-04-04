package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var arrays map[string][]int

func main() {
	arrays := make(map[string][]int)

	for {
		fmt.Print("Enter command: ")
		input := getUserInput()
		command := strings.Fields(input)
		command[0] = strings.ToLower(command[0])

		switch command[0] {
		case "load":
			loadArray(command, arrays)
			fmt.Println(arrays)
		case "save":
			saveArray(command, arrays)
		case "rand":
			random(command, arrays)
		case "concat":
			//
		case "free":
			free(command, arrays)
		case "remove":
			//
		case "copy":
			//
		case "sortnondesc":
			//
		case "sortnonasc":
			//
		case "shuffle":
			//
		case "stats":
			//
		case "print":
			//
		case "exit":
			fmt.Println("Bye :)")
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Try again")
		}
	}

}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func loadArray(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("invalid command")
		return
	}

	arrayName := command[1]
	fileName := command[2]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for i := 0; i < len(fields); i++ {
			num, err := strconv.Atoi(fields[i])
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			numbers = append(numbers, num)
		}

	}

	arrays[arrayName] = numbers
	fmt.Println("Array ", arrayName, " loaded succeessfully.")
}

func saveArray(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid command")
		return
	}

	arrayName := command[1]
	fileName := command[2]

	numbers, ok := arrays[arrayName]
	if !ok {
		fmt.Println("Array ", arrayName, " doesn't exist")
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range numbers {
		fmt.Fprintln(writer, num)
	}

	writer.Flush()

	fmt.Println("Array ", arrayName, " saved to ", fileName, " succeessfully.")
}

func random(command []string, arrays map[string][]int) {
	if len(command) != 5 {
		fmt.Println("Invalid command")
		return
	}

	arrayName := command[1]
	count, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	min, err := strconv.Atoi(command[3])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	max, err := strconv.Atoi(command[4])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	seed := rand.NewSource(time.Now().UnixNano())
	var numbers []int
	for i := 0; i < count; i++ {
		r := rand.New(seed).Int()
		randomNumber := min + r%(max-min+1)
		numbers = append(numbers, randomNumber)
		// fmt.Println(randomNumber)
	}

	arrays[arrayName] = numbers
	fmt.Printf("%v random numbers were generated and added to array \"%s\"\n", count, arrayName)
}

func concat(command []string, arrays map[string][]int) {

}

func free(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	arrayName := command[1]
	_, ok := arrays[arrayName]
	if !ok {
		fmt.Printf("Array \"%s\" does not exists\n", arrayName)
		return
	}

	arrays[arrayName] = []int{}

	fmt.Printf("Arrays \"%s\" is empty now\n", arrayName)
}
