package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"
	"strconv"
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

func selectWords(selectFunction func(word string) bool) []string {
	words := []string{}
	txtFiles, err := fs.Glob(os.DirFS("."), "*.txt")
	if err != nil {
		panic(err)
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
			if selectFunction(word) {
				words = append(words, word)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	sort.Strings(words)
	return words
}

func consistingOf(word string, letters string) bool {
	return len(word) >= 4 && hasLetter(word, []rune(letters)[0]) && !hasOtherLetters(word, []rune(letters))
}

func askLetters(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	letters := ""
	if scanner.Scan() {
		letters = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return letters
}

func solve(letters string) []string {
	selectFunction := func(word string) bool {
		return consistingOf(word, letters)
	}
	return selectWords(selectFunction)
}

func findLetters(word string) string {
	charSet := make(map[rune]bool, len(word))
	for _, char := range word {
		charSet[char] = true
	}
	letters := make([]rune, 0, len(charSet))
	for char := range charSet {
		letters = append(letters, char)
	}
	return string(letters)
}

func generate(length int) []string {
	selectFunction := func(word string) bool {
		if len(word) != length {
			return false
		}
		charSet := make(map[rune]bool, len(word))
		for _, char := range word {
			charSet[char] = true
		}
		return len(charSet) == 7
	}
	return selectWords(selectFunction)
}

type LettersScore struct {
	Letters string
	Score   int
}

func letterScores(word string) []LettersScore {
	initialLetters := []rune(findLetters(word))
	scores := make(map[string]int, len(initialLetters))
	for i := 0; i < len(initialLetters); i++ {
		letters := append(initialLetters[i:], initialLetters[0:i]...)
		firstLetter := letters[0]
		otherLetters := []rune(string(letters[1:]))
		sort.Slice(otherLetters, func(i, j int) bool {
			return otherLetters[i] < otherLetters[j]
		})
		letters = append([]rune{firstLetter}, otherLetters...)
		scores[string(letters)] = len(solve(string(letters)))
	}
	return sortMap(scores)
}

func sortMap(values map[string]int) []LettersScore {
	var pairs []LettersScore
	for k, v := range values {
		pairs = append(pairs, LettersScore{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Score > pairs[j].Score
	})
	return pairs
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s {num}|{word}|{letters}\n", os.Args[0])
		fmt.Println()
		fmt.Println("num: Number of characters in the seeding word")
		fmt.Println("word: Seeding word consisting of 7 unique characters")
		fmt.Println("letters: 7 unique characters with mandatory character first")
		fmt.Println()
		return
	}
	length, err := strconv.Atoi(os.Args[1])
	if err == nil {
		words := generate(length)
		for _, word := range words {
			fmt.Println(word)
		}
		return
	}
	if len(os.Args[1]) > 7 {
		scores := letterScores(os.Args[1])
		for _, ls := range scores {
			fmt.Printf("%s: %d\n", ls.Letters, ls.Score)
		}
		return
	}
	words := solve(os.Args[1])
	for _, word := range words {
		fmt.Println(word)
	}
}
