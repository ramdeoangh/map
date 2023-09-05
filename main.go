package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF5733",
		"blaze": "#BD472E",
		"green": "#57BD2E",
	}
	printMap(colors)

	// others way

	// var colors2 map[int]string

	// colors2 := make(map[string]int)

	// colors2["red"] = 10

	// delete(colors2, "red")
	// fmt.Println(colors2)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
