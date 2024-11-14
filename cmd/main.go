package main

import (
	"fmt"
	"time"

	"math/rand"
)

//"github.com/fatih/color"

const (
	theWord = "HELLO, WORLD!"
)

func randomPrint() rune {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ!,")
	b := chars[rnd.Intn(len(chars))]
	return b
}

func main() {
	text := []rune("HELLO, WORLD!")
	result := ""

	for i := 0; i <= len(text)-1; i++ {
		for j := 0; j <= 50; j++ {
			fmt.Print("\r" + result + string(randomPrint()))
			time.Sleep(20 * time.Millisecond)
			fmt.Print("\r                    \r")
		}
		fmt.Print("\r" + result + string(text[i]))
		result += string(text[i])

	}

}
