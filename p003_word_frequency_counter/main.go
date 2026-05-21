package main

/*
===========================================
  Go 100 Challenge
  Problem: #003
  Level: 🟢 Small
===========================================

Problem: Word Frequency Counter

Topic:
- map
- strings package
- for range loop

Industry Use:
Search Engines / Text Analytics / SEO Tools /
Log Analysis Systems

Rules:
- User একটা sentence input দেবে
- প্রতিটা word কতবার আছে count করবে
- Case-insensitive ("Go" আর "go" same word)
- Punctuation ignore করবে (. , ! ?)
- Output: word => count, frequency অনুযায়ী
  descending sort করে print করবে

Example Run:
  Enter sentence:
  Go is great. go is fast. Go!

  Output:
  go    => 3
  is    => 2
  great => 1
  fast  => 1
===========================================
*/

import (
	"fmt"
	"unicode"

	"strings"
)

func wordFrequency(sentence string) {
	lowerCaseSentence := strings.ToLower(sentence)

	// remove Punctuation method ---
	removePunct := func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}

	// Punctuation removed result ---
	punctRemovedSentence := strings.Map(removePunct, lowerCaseSentence)

	// split sentence ---
	words := strings.Fields(punctRemovedSentence)

	// fmt.Println(sentenceSplit)

	// count multiple sentence
	sentenceMap := make(map[string]int)

	for _, word := range words {
		trimmedSentence := strings.TrimSpace(word)

		if trimmedSentence != "" {

			sentenceMap[trimmedSentence]++
		}
	}

	for s, count := range sentenceMap {
		fmt.Printf("'%s' এসেছে: %d বার\n", s, count)
	}

}

func main() {
	var sentence string

	sentence = "How are you BRo!!@# . how are you"

	wordFrequency(sentence)
}
