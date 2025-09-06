package token

import "testing"

func TestMint(t *testing.T) {
	tkn := Mint("1")
	found := false
	for _, c := range Characteristics {
		if tkn.Characteristic == c {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Característica inválida: %s", tkn.Characteristic)
	}
}
