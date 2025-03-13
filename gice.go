package main

import (
	"fmt"
	"giceware/data"
	"log"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func concatNum(a, b int) int {
	out, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return out
}

func rollDie() int {
	return rand.IntN(6) + 1
}

func rollWord() int {
	var word int
	var arr []int
	for i := 0; i < 5; i++ {
		arr = append(arr, rollDie())
	}
	for i, element := range arr {
		if i == 0 {
			word = element
		} else {
			word = concatNum(word, element)
		}

	}
	return word
}

func rollPhrase(wordsNum int) []int {
	var arr []int
	for i := 0; i < wordsNum; i++ {
		arr = append(arr, rollWord())
	}
	return arr
}

func sprinkleNumbers(phrase []string, numberOfNumbers int) []string {
	for i := 0; i < numberOfNumbers; i++ {
		num := rand.IntN(10)
		position := rand.IntN(len(phrase))
		phrase[position] += strconv.Itoa(num)
	}
	return phrase
}

func generatePhrase(numberOfWOrds int, numberOfNumbers int, separator string) string {
	wordsInNumbers := rollPhrase(numberOfWOrds)
	var words []string
	for _, word := range wordsInNumbers {
		words = append(words, data.Wordlist[word])
	}
	words = sprinkleNumbers(words, numberOfNumbers)
	var phrase string
	for i, word := range words {
		if i == 0 {
			phrase = word
		} else {
			phrase = phrase + separator + word
		}
	}
	return phrase
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:    "words",
				Usage:   "set the amount of words your password should have",
				Aliases: []string{"w"},
				Value:   6,
			},
			&cli.Int64Flag{
				Name:    "numbers",
				Usage:   "set the number or numbers, please note that if the number of numbers is larger than the  number of words you will get multi digit numbers",
				Aliases: []string{"n"},
				Value:   1,
			},
			&cli.StringFlag{
				Name:    "separator",
				Usage:   "Set the separation character",
				Aliases: []string{"s"},
				Value:   "-",
			},
		},
		Action: func(ctx *cli.Context) error {
			words := ctx.Int64("words")
			if words <= 0 {
				return fmt.Errorf("can't set 0 or less words")
			}
			numbers := ctx.Int64("numbers")
			if numbers < 0 {
				return fmt.Errorf("cant set negative amount of numbers")
			}
			seperator := ctx.String("separator")
			phrase := generatePhrase(int(words), int(numbers), seperator)
			fmt.Printf("Your passphrase is:\n%v\n\n", phrase)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
