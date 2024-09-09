package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("teste.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := bufio.NewReader(os.Stdin)
	i := 0
	for {
		fmt.Printf("%d >", i) // "ask" for input
		i++
		line, _, err := input.ReadLine() // awaits the input
		if err != nil {
			panic(err)
		}

		f.Write([]byte("\n")) // write in the file itself
		f.Write(line)
	}
}
