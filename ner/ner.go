package ner

import (
	"fmt"
	"log"

	prose "gopkg.in/jdkato/prose.v2"
)

//TestNer prints out tokens from a sample string
func TestNer(input string) {
	doc, err := prose.NewDocument(input)
	if err != nil {
		log.Fatal(err)
	}
	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
	}
}
