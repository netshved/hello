package random

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/fatih/color"
)

func RandomColor() *color.Color {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	colors := []*color.Color{
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgBlue),
		color.New(color.FgMagenta),
		color.New(color.FgCyan),
		color.New(color.FgWhite),
		color.New(color.FgHiRed),
		color.New(color.FgHiGreen),
		color.New(color.FgHiYellow),
		color.New(color.FgHiBlue),
		color.New(color.FgHiMagenta),
		color.New(color.FgHiCyan),
		color.New(color.FgHiWhite),
	}

	b := colors[rnd.Intn(len(colors))]

	return b

}

func RandomChar() rune {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ!,")
	b := chars[rnd.Intn(len(chars))]
	return b
}

func RandomPrint(text []rune) {
	result := ""
	for i := 0; i <= len(text)-1; i++ {
		for j := 0; j <= 30; j++ {
			part1 := result + string(RandomChar())
			RandomColor().Print("\r" + part1)
			time.Sleep(20 * time.Millisecond)
			fmt.Print("\r                    \r")
		}

		RandomColor().Print("\r" + result + string(text[i]))
		result += string(text[i])

	}
}
