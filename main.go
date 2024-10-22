package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	// Read words from a file
	wordFile := "dict.txt"
	words, err := readWords(wordFile)
	if err != nil {
		fmt.Println("Error reading wordset:", err)
		return
	}

	// Shuffle the words randomly
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	usedWords := make(map[string]bool) // Track used words

	scanner := bufio.NewScanner(os.Stdin)
	var lastPhrase string // Store the last input phrase
	fmt.Println("Enter a phrase to search (or type 'exit' to quit):")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		phrase := scanner.Text()
		if strings.ToLower(phrase) == "exit" {
			break
		} else if phrase == "" {
			// If the input is empty, use the last phrase
			phrase = lastPhrase
		} else {
			// Update the last phrase if the current input is not empty
			lastPhrase = phrase
		}

		// Find the shortest matching word that hasn't been used yet
		shortestWord := findShortestWord(words, strings.ToUpper(phrase), usedWords)
		if shortestWord != "" {
			fmt.Println("Shortest word:", shortestWord)
			// Copy to clipboard
			clipboard.WriteAll(shortestWord)
			usedWords[shortestWord] = true
		} else {
			fmt.Println("No word contains the phrase:", phrase)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

// Function to read words from a file into a slice
func readWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text()) // Ensure no extra whitespace
		if word != "" {
			words = append(words, word)
		}
	}

	return words, scanner.Err()
}

// Function to find the shortest word that contains the phrase and hasn't been used yet
func findShortestWord(words []string, phrase string, usedWords map[string]bool) string {
	phrase = strings.TrimSpace(phrase) // Clean the input phrase
	var shortestWord string
	for _, word := range words {
		if strings.Contains(word, phrase) && !usedWords[word] { // Ignore used words
			if shortestWord == "" || len(word) < len(shortestWord) {
				shortestWord = word
			}
		}
	}
	return shortestWord
}
