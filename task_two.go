package main

import (
	"fmt"
	"strings"
)

// Word frequency counter
func WordFrequency(text string) map[string]int {
	frequency := make(map[string]int)

	text = strings.ToLower(text)
	cleaned := ""

	for _, r := range text {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == ' ' {
			cleaned += string(r)
		}
	}

	words := strings.Fields(cleaned)
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

// Palindrome checker
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	cleaned := ""

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			cleaned += string(r)
		}
	}

	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	text := "Hello, hello! How are you? Are you okay, hello?"
	fmt.Println("Word Frequencies:")
	for word, count := range WordFrequency(text) {
		fmt.Printf("%s: %d\n", word, count)
	}

	fmt.Println()

	sample := "A man, a plan, a canal: Panama"
	if IsPalindrome(sample) {
		fmt.Printf("'%s' is a palindrome.\n", sample)
	} else {
		fmt.Printf("'%s' is not a palindrome.\n", sample)
	}
}
