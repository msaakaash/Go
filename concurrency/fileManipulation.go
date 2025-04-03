package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func writeFile(file string, messages []string) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer f.Close()

	for _, message := range messages {
		_, err = f.WriteString(message + "\n")
		if err != nil {
			fmt.Println("Error writing to file")
			return
		}
	}
}

func readFile(file string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ch <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
	}

	close(ch)
}

func processFile(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for line := range ch {
		fmt.Println("Processed line:", line)
	}
}

func main() {
	var inputLines []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text to write to file (type 'STOP' on a new line to finish):")

	for {
		text, _ := reader.ReadString('\n')
		if text == "STOP" {
			break
		}
		inputLines = append(inputLines, text)
	}

	filename := "sample.txt"
	writeFile(filename, inputLines)

	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go readFile(filename, ch, &wg)

	wg.Add(1)
	go processFile(ch, &wg)

	wg.Wait()
}
