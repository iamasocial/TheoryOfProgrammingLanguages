package funcs

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ArrayExists(name string, arrays map[string][]int) bool {
	_, ok := arrays[name]
	if !ok {
		fmt.Printf("Array \"%s\" does not exists\n", name)
		return ok
	}

	return ok
}

func QuickSort(array []int, sign string) []int {
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
		return append(append(QuickSort(less, sign), op), QuickSort(more, sign)...)
	} else {
		return append(append(QuickSort(more, sign), op), QuickSort(less, sign)...)
	}
}

func Mix(arr []int) []int {
	seed := rand.NewSource(time.Now().UnixNano())
	rand.New(seed)
	length := len(arr) - 1
	for i := length; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
