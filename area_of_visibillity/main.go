package main

import (
	"areaOfVisibillity/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("aov.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	currentArea := utils.NewArea(nil)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "{":
			currentArea = utils.NewArea(currentArea)
		case "}":
			currentArea = currentArea.Parent
		case "ShowVar;":
			for key, value := range currentArea.Vars {
				fmt.Printf("%s = %v\n", key, value)
			}
		default:
			if strings.Contains(fields[0], "var") && strings.Contains(fields[0], "=") && strings.Contains(fields[0], ";") {
				index := strings.Index(fields[0], "=")
				name := fields[0][:index]
				value, err := strconv.Atoi(fields[0][index+1 : len(fields[0])-1])
				if err != nil {
					panic("Error!")
				}
				currentArea.Vars[name] = value
			} else {
				panic("Govno, ne rabotaet")
			}
		}
	}
}
