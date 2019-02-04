package main

import "fmt"

func main() {
	maps := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	fmt.Println(maps)
	printMap(maps)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("For Color: ", color, "Hex is:", hex)
	}
}
