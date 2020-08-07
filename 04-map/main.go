package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(colors map[string]string) {
	for key, value := range colors {
		fmt.Println("Hex code for", key, "is", value)
	}
}

// colors := make(map[string]string)
// colors["white"] = "#ffffff"
// delete(colors, "white")

// var colors map[string]string
// colors["white"] = "#ffffff"

// colors := map[string]string{
// 	"red":   "#ff0000",
// 	"green": "#4bf745",
// }
