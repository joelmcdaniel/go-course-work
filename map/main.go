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

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

/* func main() {
	m := map[string]string{
		"dog": "bark",
	}

	changeMap(m)

	fmt.Println(m)
}

func changeMap(m map[string]string) {
	m["cat"] = "purr"
} */
