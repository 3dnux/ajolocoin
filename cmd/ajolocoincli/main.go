package main

import (
	"flag"
	"fmt"

	"ajolocoin/token"
)

func main() {
	id := flag.String("id", "1", "ID del token")
	flag.Parse()
	t := token.Mint(*id)
	fmt.Printf("Token %s con característica %s\n", t.ID, t.Characteristic)
}
