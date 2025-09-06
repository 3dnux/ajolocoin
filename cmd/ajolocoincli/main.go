package main

import (
	"fmt"
	"os"

	"ajolocoin/app"
)

// Simple CLI to mint an AjoloteToken with a characteristic.
func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: ajolocoincli <id> <characteristic>")
		os.Exit(1)
	}

	id := os.Args[1]
	characteristic := os.Args[2]

	a := app.Setup()
	token := a.Mint(id, characteristic)
	fmt.Printf("Minted token %s with characteristic %s\n", token.ID, token.Characteristic)
}
