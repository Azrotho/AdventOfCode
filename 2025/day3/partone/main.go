package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	total := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		wherefirstdigit := getMaxInStringFD(text)
		suite := SplitAtChar(text, wherefirstdigit)
		whereisseconddigit := getMaxInStringSD(suite)
		result := string(text[wherefirstdigit]) + string(suite[whereisseconddigit])
		fmt.Printf("%s\n", result)
		val, err := strconv.Atoi(result)
		if err == nil {
			total += val
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", total)
}

func getMaxInStringFD(chaine string) int {
	actualmax := -1
	wheremax := -1
	for i := range len(chaine) - 1 {
		val, err := strconv.Atoi(string(chaine[i]))
		if err == nil && val > actualmax {
			actualmax = val
			wheremax = i
		}
	}
	return wheremax
}

func getMaxInStringSD(chaine string) int {
	actualmax := -1
	wheremax := -1
	for i := range len(chaine) {
		val, err := strconv.Atoi(string(chaine[i]))
		if err == nil && val > actualmax {
			actualmax = val
			wheremax = i
		}
	}
	return wheremax
}

func SplitAtChar(str string, n int) string {
	runes := []rune(str)
	return string(runes[n+1:])
}
