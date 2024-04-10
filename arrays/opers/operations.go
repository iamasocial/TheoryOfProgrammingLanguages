package opers

import (
	"arrays/funcs"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func LoadArray(command []string, arrays map[string][]int) {
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
	fmt.Println(arrays[arrayName])
}

func SaveArray(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid command")
		return
	}

	arrayName := command[1]
	fileName := command[2]

	if !funcs.ArrayExists(arrayName, arrays) {
		return
	}
	numbers := arrays[arrayName]

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

func Random(command []string, arrays map[string][]int) {
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

func Concat(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid format")
		return
	}

	arrayName1 := command[1]
	arrayName2 := command[2]

	if !funcs.ArrayExists(arrayName1, arrays) {
		return
	}

	if !funcs.ArrayExists(arrayName2, arrays) {
		return
	}

	length := len(arrays[arrayName2])
	for i := 0; i < length; i++ {
		arrays[arrayName1] = append(arrays[arrayName1], arrays[arrayName2][i])
	}

	fmt.Println("Arrays were concatenated successfully.")
}

func Free(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	arrayName := command[1]
	if !funcs.ArrayExists(arrayName, arrays) {
		return
	}

	arrays[arrayName] = []int{}

	fmt.Printf("Arrays \"%s\" is empty now\n", arrayName)
}

func Remove(command []string, arrays map[string][]int) {
	if len(command) != 4 {
		fmt.Println("Invalid command")
		return
	}

	arrayName := command[1]
	if !funcs.ArrayExists(arrayName, arrays) {
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

func Copy(command []string, arrays map[string][]int) {
	if len(command) != 5 {
		fmt.Println("Invalid command.")
		return
	}

	arrayName1 := command[1]
	arrayName2 := command[4]

	if !funcs.ArrayExists(arrayName1, arrays) {
		return
	}

	if !funcs.ArrayExists(arrayName2, arrays) {
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

func Sort(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid Command")
		return
	}

	arrayName := command[1]
	if !funcs.ArrayExists(arrayName, arrays) {
		return
	}

	sign := command[2]
	arrays[arrayName] = funcs.QuickSort(arrays[arrayName], sign)
	fmt.Println("Sorted array: ", arrays[arrayName])
}

func Shuffle(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	arrayName := command[1]

	if !funcs.ArrayExists(arrayName, arrays) {
		return
	}

	// fmt.Println(arrays[arrayName])

	arrays[arrayName] = funcs.Mix(arrays[arrayName])
	fmt.Println(arrays[arrayName])
	fmt.Println("Array was shuffled succeessfully")
}
