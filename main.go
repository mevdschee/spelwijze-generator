package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

func selectWords(txtFile string, selectFunction func(word string) bool) []string {
	words := []string{}
	file, err := os.Open(txtFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	gz, err := gzip.NewReader(file)
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	for scanner.Scan() {
		word := scanner.Text()
		if selectFunction == nil || selectFunction(word) {
			words = append(words, word)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
	return selectWords("words.txt.gz", selectFunction)
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

func findWordFrequency() map[string]int {
	allWords := selectWords("words.txt.gz", nil)
	wordFrequency := make(map[string]int, len(allWords))
	for _, word := range allWords {
		wordFrequency[word] = 0
	}
	setFrequency := func(wordAndFrequency string) bool {
		parts := strings.Split(wordAndFrequency, "\t")
		word := parts[0]
		frequency, err := strconv.Atoi(parts[1])
		if err == nil {
			_, ok := wordFrequency[word]
			if ok {
				wordFrequency[word] = frequency
			}
		}
		return false
	}
	selectWords("wordfreq.txt.gz", setFrequency)
	return wordFrequency
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
	return selectWords("words.txt.gz", selectFunction)
}

type LettersScore struct {
	Letters string
	Score   int
}

func letterScores(word string) []LettersScore {
	wordFrequency := findWordFrequency()
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
		words := solve(string(letters))
		count := len(words)
		score := 0
		for _, word := range words {
			if wordFrequency[word] > 0 {
				score += len(word) + wordFrequency[word]/10
			}
			if wordFrequency[word] == 0 {
				score -= 30
			}
		}
		scores[string(letters)] = count*2 + score
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
		fmt.Printf("Usage: %s {number}|{word}|{letters}\n", os.Args[0])
		fmt.Println()
		fmt.Println("number  : Number of letters in the seeding word")
		fmt.Println("word    : Seeding word consisting of 7 unique letters")
		fmt.Println("letters : 7 unique letters with mandatory letter first")
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
