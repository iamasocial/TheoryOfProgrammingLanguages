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
			concat(command, arrays)
		case "free":
			free(command, arrays)
		case "remove":
			remove(command, arrays)
		case "copy":
			copy(command, arrays)
		case "sort":
			sort(command, arrays)
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

	numbers, _ := arrays[arrayName]
	// if !ok {
	// 	fmt.Println("Array ", arrayName, " doesn't exist")
	// 	return
	// }
	if !arrayExists(arrayName, arrays) {
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
		fmt.Println(randomNumber)
	}

	arrays[arrayName] = numbers
	fmt.Printf("%v random numbers were generated and added to array \"%s\" successfully.\n", count, arrayName)
}

func concat(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid format")
		return
	}

	arrayName1 := command[1]
	arrayName2 := command[2]

	if !arrayExists(arrayName1, arrays) {
		return
	}

	if !arrayExists(arrayName2, arrays) {
		return
	}

	length := len(arrays[arrayName2])
	for i := 0; i < length; i++ {
		arrays[arrayName1] = append(arrays[arrayName1], arrays[arrayName2][i])
	}

	fmt.Println("Arrays were concatenated successfully.")
}

func free(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	arrayName := command[1]
	if !arrayExists(arrayName, arrays) {
		return
	}

	arrays[arrayName] = []int{}

	fmt.Printf("Arrays \"%s\" is empty now\n", arrayName)
}

func remove(command []string, arrays map[string][]int) {
	if len(command) != 4 {
		fmt.Println("Invalid command")
		return
	}

	arrayName := command[1]
	if !arrayExists(arrayName, arrays) {
		return
	}

	index, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Println("Bad index")
		return
	}

	count, err := strconv.Atoi(command[3])
	if err != nil {
		fmt.Println("Bad count")
		return
	}

	if index+count > len(arrays[arrayName]) {
		fmt.Println("You want to remove too much. Try again")
		return
	}

	tmp := arrays[arrayName][:index]
	tmp = append(tmp, arrays[arrayName][count+index:]...)
	arrays[arrayName] = tmp
	fmt.Println(arrays[arrayName])
	fmt.Println("Elements were removed successfully.")
}

func copy(command []string, arrays map[string][]int) {
	if len(command) != 5 {
		fmt.Println("Invalid command.")
		return
	}

	arrayName1 := command[1]
	arrayName2 := command[4]

	if !arrayExists(arrayName1, arrays) {
		return
	}

	if !arrayExists(arrayName2, arrays) {
		return
	}

	begin, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Println("First index is bad.")
		return
	}

	end, err := strconv.Atoi(command[3])
	if err != nil {
		fmt.Println("Second index is bad.")
		return
	}

	if len(arrays[arrayName1]) < begin || len(arrays[arrayName2]) < end {
		fmt.Println("Index is out of range.")
		return
	}

	arrays[arrayName2] = append(arrays[arrayName2], arrays[arrayName1][begin:end+1]...)
	fmt.Println(arrays[arrayName2])
	fmt.Println("Copied successfully.")
}

func sort(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid Command")
		return
	}

	arrayName := command[1]
	if !arrayExists(arrayName, arrays) {
		return
	}

	sign := command[2]
	arrays[arrayName] = quickSort(arrays[arrayName], sign)
	fmt.Println("Sorted array: ", arrays[arrayName])
}

func shuffle(command []string, arrays map[string][]int) {

}

func arrayExists(name string, arrays map[string][]int) bool {
	_, ok := arrays[name]
	if !ok {
		fmt.Printf("Array \"%s\" does not exists\n", name)
		return ok
	}

	return ok
}

func quickSort(array []int, sign string) []int {
	if len(array) < 2 {
		return array
	}

	if sign != "+" && sign != "-" {
		fmt.Println("Bad sort argument")
		return array
	}
	var less, more []int

	op := array[0]
	for _, value := range array[1:] {
		if value <= op {
			less = append(less, value)
		} else {
			more = append(more, value)
		}
	}

	if sign == "+" {
		return append(append(quickSort(less, sign), op), quickSort(more, sign)...)
	} else {
		return append(append(quickSort(more, sign), op), quickSort(less, sign)...)
	}
}
