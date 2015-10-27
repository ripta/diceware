package main

import "crypto/rand"
import "fmt"
import "log"
import "math/big"
import "strconv"

import "github.com/codegangsta/cli"

func GeneratePassphrase(c *cli.Context) {
	wordlistName := c.GlobalString("word-list")
	numWords     := c.GlobalInt("num-words")
	showNumbers  := !c.GlobalBool("quiet")
	debugMode    := false

	words, err := getWordlistByName(wordlistName)
	if err != nil {
		log.Fatalf("Could not load word list '%s': %q\n", wordlistName, err)
	}

	if debugMode {
		log.Printf("Loaded %d words\n", len(words))
	}

	rolls, err := rollSets(numWords, len(words))
	if err != nil {
		log.Fatalf("Error when randomizing: %s\n", err)
	}

	if debugMode {
		log.Printf("Rolls: %v\n", rolls)
	}

	chosenWords := make([]string, 0, numWords)
	for _, index := range rolls {
		chosenWords = append(chosenWords, words[index])
	}

	if showNumbers {
		for _, roll := range rolls {
			fmt.Printf("%s ", strconv.FormatInt(roll, 10))
		}
		fmt.Println()
	}

	for _, chosenWord := range chosenWords {
		fmt.Printf("%s ", chosenWord)
	}
	fmt.Println()
}

func rollDice(max *big.Int) (randomValue *big.Int, err error) {
	return rand.Int(rand.Reader, max)
}

func rollSets(numRolls, max int) (rolls []int64, err error) {
	for i := 0; i < numRolls; i++ {
		roll, rollErr := rollDice(big.NewInt(int64(max)))
		if rollErr != nil {
			err = fmt.Errorf("Error while rolling dice: %s", rollErr)
			return
		}

		rolls = append(rolls, roll.Int64())
	}

	return
}
