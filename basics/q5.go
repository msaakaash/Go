package main

import (
	"fmt"
	"os"
)

func main() {
	// List of cities
	cities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix"}

	// Create or open a file to write to
	file, err := os.Create("cities.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the cities to the file
	for _, city := range cities {
		_, err := file.WriteString(city + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Cities have been written to cities.txt")
}
