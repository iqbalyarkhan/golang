package main

import (
	"bufio"
	"fmt"
	"ner"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to tokenize: ")
	text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	ner.TestNer(text)
}
