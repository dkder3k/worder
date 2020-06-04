package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	dictionaryFilePath := flag.String("dictionary", "", "Path to a dictionary file")
	flag.Parse()

	word := os.Args[2]

	file, err := os.Open(*dictionaryFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bestGuessWord := ""
	bestSubsequencesLength := 0
	bestRunesMatch := 0

	guessWords := make(map[string][2]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dictWord := scanner.Text()
		if len(dictWord) > len(word) || len(dictWord) < len(word) {
			continue
		}

		d := calculateSubsequenceLength(word, dictWord)
		if d < len(word) - 2 {
			continue
		}
		if len(word) == d && len(word) == len(dictWord) {
			bestGuessWord = dictWord
			fmt.Println("That's a word!")
			break
		}

		m := 0
		for i := 0; i < len(dictWord); i++ {
			if word[i] == dictWord[i] {
				m++
			}
		}

		guessWords[dictWord] = [2]int{d, m}
	}

	if bestGuessWord == "" {
		for w, v := range guessWords {
			if v[0] > bestSubsequencesLength && v[1] > bestRunesMatch {
				bestGuessWord = w
				bestSubsequencesLength = v[0]
				bestRunesMatch = v[1]
			}
		}
	}

	fmt.Printf("Did you mean %v?\n", bestGuessWord)

}

func calculateSubsequenceLength(word1, word2 string) int {
	matrix := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		matrix[i] = make([]int, len(word2))
	}

	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			leftNeighbor, upperNeighbor := 0, 0
			if i > 0 {
				upperNeighbor = matrix[i - 1][j]
			}
			if j > 0 {
				leftNeighbor = matrix[i][j - 1]
			}

			if leftNeighbor > upperNeighbor {
				matrix[i][j] = leftNeighbor
			} else {
				matrix[i][j] = upperNeighbor
			}

			if word1[i] == word2[j] {
				if i > 0 && j > 0 {
					matrix[i][j] = matrix[i - 1][j - 1] + 1
				} else {
					matrix[i][j] = 1
				}
			}
		}
	}


	return matrix[len(word1) - 1][len(word2) - 1]
}
