package main

import "fmt"

func main() {
	myData := []int{4, 6, 5, 2, 8}
	printGraph(myData)
	sortInc(myData)
	sortDec(myData)
}

func printGraph(data []int) {
	maxValue := 0
	for _, val := range data {
		if maxValue < val {
			maxValue = val
		}
	}

	hData := maxValue

	for h := 0; h < maxValue; h++ {
		for i := 0; i < len(data); i++ {
			if data[i] >= hData {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		hData--
		fmt.Println("")
	}

	for _, val := range data {
		fmt.Print(val, " ")
	}
	fmt.Println("\n ")
}

func sortInc(data []int) {
	var temp int
	done := false

	fmt.Println("=== Increase Sorting ===")
	for done == false {
		done = true
		for i := 0; i < len(data)-1; i++ {
			if data[i] > data[i+1] {
				temp = data[i]
				data[i] = data[i+1]
				data[i+1] = temp
				done = false
				printGraph(data)
			}
		}
	}
}

func sortDec(data []int) {
	var temp int
	done := false

	println("=== Decrease Sorting ===")
	for done == false {
		done = true
		for i := 0; i < len(data)-1; i++ {
			if data[i] < data[i+1] {
				temp = data[i]
				data[i] = data[i+1]
				data[i+1] = temp
				done = false
				printGraph(data)
			}
		}
	}
}
