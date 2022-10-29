package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

var wordBank = [100]string{"influence", "combine", "word", "gave", "table", "triangle", "go", "us", "lion", "attack", "moon", "front", "independent", "ants", "point", "with", "none", "percent", "life", "pack", "doctor", "proud", "flower", "topic", "remember", "least", "electric", "together", "worker", "raise", "thread", "contrast", "dish", "pay", "glass", "ten", "although", "tall", "struck", "kill", "doubt", "brown", "union", "breeze", "chest", "safe", "till", "familiar", "river", "stems", "stove", "short", "be", "early", "simply", "shot", "thin", "take", "related", "become", "me", "each", "tobacco", "coast", "pick", "guide", "shut", "lie", "exchange", "throat", "additional", "tales", "fun", "forward", "trip", "pet", "fuel", "especially", "start", "stiff", "goes", "angry", "range", "capital", "meant", "supper", "swam", "report", "fellow", "still", "broken", "tower", "excited", "handsome", "having", "industrial", "base"}

func main() {

	solutionIndex := rand.Intn(100)
	fmt.Println(solutionIndex)
	solution := wordBank[solutionIndex]
	guessLimit := len(solution) + 1
	solveKey := generateSolveKey(solution)
	// fmt.Println(solution, guessCount, guessLimit)
	printBoard(guessLimit, len(solution))
	resultBoard := ""
	gameWon := false

	for guessCount := 0; guessCount < guessLimit; guessCount++ {
		guess := ""
		guess = acceptGuess()
		result := checkGuess(guess, solution)
		resultBoard += (result + "\n")
		fmt.Print(resultBoard)
		if result == solveKey {
			fmt.Println("WON")
			gameWon = true
			guessCount = guessLimit
		}

	}

	if gameWon == false {
		fmt.Println("Solution: " + solution)
	}

}

func generateSolveKey(solution string) string {
	key := ""
	for i := 0; i < len(solution); i++ {
		key += "✅"
	}
	return key
}

func printBoard(totalGuesses, wordLength int) {
	for i := 0; i < totalGuesses; i++ {
		line := ""
		for j := 0; j < wordLength; j++ {
			line += "▮"
		}
		fmt.Println(line)
	}
}

func acceptGuess() string {

	fmt.Println("please enter your guess:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return scanner.Text()

}

func checkGuess(guess, solution string) string {
	result := ""
	for i := 0; i < (len(guess)); i++ {
		if guess[i] == solution[i] {
			result += "✅"
		} else if strings.Contains(solution, string(guess[i])) {
			result += "🟨"
		} else {
			result += "❌"
		}
	}
	return result
}
