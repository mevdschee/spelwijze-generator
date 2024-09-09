package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"
)

func hasLetter(s string, l rune) bool {
	for _, r := range s {
		if r == l {
			return true
		}
	}
	return false
}

func hasOtherLetters(s string, letters []rune) bool {
	for _, r := range s {
		if !hasLetter(string(letters), r) {
			return true
		}
	}
	return false
}

func main() {
	words := []string{}
	txtFiles, err := fs.Glob(os.DirFS("."), "*.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter character (mandatory first): ")
	letters := ""
	if scanner.Scan() {
		letters = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, txtFile := range txtFiles {
		file, err := os.Open(txtFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := scanner.Text()
			if len(word) >= 4 && hasLetter(word, []rune(letters)[0]) && !hasOtherLetters(word, []rune(letters)) {
				words = append(words, word)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	sort.Strings(words)
	for _, word := range words {
		fmt.Println(word)
	}
}
