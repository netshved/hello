package main

import (
	"bufio"
	"os"
	"time"

	r "github.com/netshved/hello/pkg/random"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	r.RandomColor().Print("Print word to know the truly color: ")
	text, _ := reader.ReadString('\n')

	r.RandomPrint([]rune(text))
	time.Sleep(120 * time.Second)
}
