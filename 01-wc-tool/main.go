package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 3 {
		fmt.Println("Usage: ./cwcc [option] <file>")
		return
	}

	var fileName string
	var flag string

	lines, words, characters := 0, 0, 0

	for i := 0; i < len(os.Args); i++ {
		arg := os.Args[i]

		if len(arg) > 0 && arg[0] == '-' {
			flag = arg[1:]
		} else {
			fileName = arg
		}
	}

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Printf("Error getting file information %s\n", err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		characters += len(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	bytes := fileInfo.Size()

	switch flag {
	case "c":
		fmt.Println(bytes, fileName)
		break
	case "l":
		fmt.Println(lines, fileName)
		break
	case "w":
		fmt.Println(words, fileName)
		break
	case "m":
		fmt.Println(characters, fileName)
		break
	default:
		fmt.Println(lines, words, bytes, fileName)
	}
}
