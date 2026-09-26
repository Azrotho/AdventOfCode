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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		currentSuite := text
		result := ""
		validLine := true
		for n := 11; n >= 0; n-- {
			whereDigit := getMaxInString(currentSuite, n)
			if whereDigit == -1 {
				validLine = false
				break
			}
			char := string(currentSuite[whereDigit])
			result += char
			if n > 0 {
				currentSuite = SplitAtChar(currentSuite, whereDigit)
			}
		}

		if validLine {
			val, err := strconv.Atoi(result)
			if err == nil {
				total += val
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", total)
}

func getMaxInString(chaine string, n int) int {
	actualmax := -1
	wheremax := -1
	for i := range len(chaine) - n {
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
