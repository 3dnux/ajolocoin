package token

import (
	"math/rand"
	"time"
)

type Token struct {
	ID             string
	Characteristic string
}

var Characteristics = []string{
	"Albino",
	"Leucístico",
	"Melanoide",
	"Dorado",
	"Salvaje",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Mint(id string) Token {
	trait := Characteristics[rand.Intn(len(Characteristics))]
	return Token{ID: id, Characteristic: trait}
}
