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

type Func func([]string, map[string][]int)

var FuncMap map[string]Func
var Arrays map[string][]int

func LoadArray(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("invalid command")
		return
	}

	name := command[1]
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

	arrays[name] = numbers
	fmt.Println("Array ", name, " loaded succeessfully.")
}

func SaveArray(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]
	fileName := command[2]

	if !funcs.ArrayExists(name, arrays) {
		return
	}
	numbers := arrays[name]

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

	fmt.Println("Array ", name, " saved to ", fileName, " succeessfully.")
}

func Random(command []string, arrays map[string][]int) {
	if len(command) != 5 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]
	count, err := strconv.Atoi(command[2])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	firstValue, err := strconv.Atoi(command[3])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	secondValue, err := strconv.Atoi(command[4])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var min, max int
	if firstValue > secondValue {
		max = firstValue
		min = secondValue
	} else {
		max = secondValue
		min = firstValue
	}

	seed := rand.NewSource(time.Now().UnixNano())
	var numbers []int
	for i := 0; i < count; i++ {
		r := rand.New(seed).Int()
		randomNumber := min + r%(max-min+1)
		numbers = append(numbers, randomNumber)
	}

	arrays[name] = numbers
	fmt.Printf("%v random numbers were generated and added to array \"%s\" successfully.\n", count, name)
}

func Concat(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid format")
		return
	}

	name1 := command[1]
	name2 := command[2]

	if !funcs.ArrayExists(name1, arrays) {
		return
	}

	if !funcs.ArrayExists(name2, arrays) {
		return
	}

	length := len(arrays[name2])
	for i := 0; i < length; i++ {
		arrays[name1] = append(arrays[name1], arrays[name2][i])
	}

	fmt.Println("Arrays were concatenated successfully.")
}

func Free(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]
	if !funcs.ArrayExists(name, arrays) {
		return
	}

	arrays[name] = []int{}

	fmt.Printf("Arrays \"%s\" is empty now\n", name)
}

func Remove(command []string, arrays map[string][]int) {
	if len(command) != 4 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]
	if !funcs.ArrayExists(name, arrays) {
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

	if index+count > len(arrays[name]) {
		fmt.Println("You want to remove too much. Try again")
		return
	}

	tmp := arrays[name][:index]
	tmp = append(tmp, arrays[name][count+index:]...)
	arrays[name] = tmp
	fmt.Println("Elements were removed successfully.")
}

func Copy(command []string, arrays map[string][]int) {
	if len(command) != 5 {
		fmt.Println("Invalid command.")
		return
	}

	name1 := command[1]
	name2 := command[4]

	if !funcs.ArrayExists(name1, arrays) {
		return
	}

	if !funcs.ArrayExists(name2, arrays) {
		return
	}

	begin, err := strconv.Atoi(command[2])
	if err != nil || begin < 0 {
		fmt.Println("First index is bad.")
		return
	}

	end, err := strconv.Atoi(command[3])
	if err != nil || end < 0 {
		fmt.Println("Second index is bad.")
		return
	}

	if len(arrays[name1])-1 < begin || len(arrays[name1])-1 < end {
		fmt.Println("Index out of range.")
		return
	}

	arrays[name2] = append(arrays[name2], arrays[name1][begin:end+1]...)
	fmt.Println("Copied successfully.")
}

func Sort(command []string, arrays map[string][]int) {
	if len(command) != 3 {
		fmt.Println("Invalid Command")
		return
	}

	name := command[1]
	if !funcs.ArrayExists(name, arrays) {
		return
	}

	sign := command[2]
	arrays[name] = funcs.QuickSort(arrays[name], sign)
	fmt.Printf("Arrays \"%s\" was sorted successfully\n", name)
}

func Shuffle(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]

	if !funcs.ArrayExists(name, arrays) {
		return
	}

	arrays[name] = funcs.Mix(arrays[name])
	fmt.Println("Array was shuffled succeessfully")
}

func Stats(command []string, arrays map[string][]int) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]

	if !funcs.ArrayExists(name, arrays) {
		return
	}

	array := arrays[name]

	length := len(array)
	max, maxIndex := funcs.GetMax(array)
	min, minIndex := funcs.GetMin(array)
	mostCommon := funcs.GetMostCommon(array)
	mean := funcs.GetMean(array)
	deviation := funcs.GetMaxDeviation(array)

	fmt.Printf("Length: %v\n", length)
	fmt.Printf("Maximum: %v (index %v)\n", max, maxIndex)
	fmt.Printf("Minimum: %v (index %v)\n", min, minIndex)
	fmt.Printf("Most common: %v\n", mostCommon)
	fmt.Printf("Mean: %v\n", mean)
	fmt.Printf("Max deviation: %v\n", deviation)
}

func Print(command []string, arrays map[string][]int) {
	length := len(command)
	if length < 3 || length > 4 {
		fmt.Println("Invalid command")
		return
	}

	name := command[1]

	if !funcs.ArrayExists(name, arrays) {
		return
	}

	if length == 3 {
		if strings.ToLower(command[2]) == "all" {
			if len(arrays[name]) == 0 {
				fmt.Printf("Array \"%s\" is empty\n", name)
				return
			}

			fmt.Println(arrays[name])
			return
		}

		index, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Println("Invalid command")
			return
		}
		if index >= len(arrays[name]) {
			fmt.Println("Index is out of range")
			return
		}
		fmt.Println(arrays[name][index])
		return
	} else if length == 4 {
		firstIndex, err := strconv.Atoi(command[2])
		if err != nil {
			fmt.Println("Bad first index")
			return
		}

		secondIndex, err := strconv.Atoi(command[3])
		if err != nil {
			fmt.Println("Bad second index")
			return
		}

		if firstIndex > secondIndex {
			fmt.Println("Second index is lesser than first")
			return
		}
		length = len(arrays[name]) - 1
		if firstIndex > length || secondIndex > length {
			fmt.Println("First and/or second index out of range")
			return
		}

		fmt.Println(arrays[name][firstIndex : secondIndex+1])
		return
	} else {
		fmt.Println("Invalid command")
		return
	}

}

func Exit(command []string, arrays map[string][]int) {
	fmt.Println("Bye :)")
	os.Exit(0)
}

func init() {
	Arrays = make(map[string][]int)
	FuncMap = map[string]Func{
		"load":    LoadArray,
		"save":    SaveArray,
		"rand":    Random,
		"concat":  Concat,
		"free":    Free,
		"remove":  Remove,
		"copy":    Copy,
		"sort":    Sort,
		"shuffle": Shuffle,
		"stats":   Stats,
		"print":   Print,
		"exit":    Exit,
	}
}
