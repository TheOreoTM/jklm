package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

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

	usedWords := make(map[string]bool) // Track used words

	scanner := bufio.NewScanner(os.Stdin)
	var lastPhrase string // Store the last input phrase
	var totalTimeTaken int
	var timesRan int
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

		// Start the timer for performance measurement
		start := time.Now()

		// Find the shortest matching word that contains the phrase
		shortestWord := findShortestContainingWord(words, strings.ToUpper(phrase), usedWords)

		// Calculate the elapsed time
		elapsed := time.Since(start)
		totalTimeTaken += int(elapsed)
		timesRan++

		if shortestWord != "" {
			fmt.Println("Shortest word:", shortestWord)
			fmt.Printf("Time taken: %s | Average time taken: %s\n", elapsed, time.Duration(totalTimeTaken/timesRan))
			// Copy to clipboard
			clipboard.WriteAll(shortestWord)
			usedWords[shortestWord] = true
		} else {
			fmt.Println("No word contains the phrase:", phrase)
			fmt.Printf("Time taken: %s\n", elapsed) // Print the time taken even if no word is found
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
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	// Shuffle the words randomly
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	// Sort words by length (ascending order)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	return words, scanner.Err()
}

// Function to find the shortest word that contains the phrase and hasn't been used yet
func findShortestContainingWord(words []string, phrase string, usedWords map[string]bool) string {
	phrase = strings.TrimSpace(phrase) // Clean the input phrase
	for _, word := range words {
		if strings.Contains(word, phrase) && !usedWords[word] { // Ignore used words
			return word
		}
	}
	return ""
}
