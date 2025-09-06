package app

import sdk "github.com/cosmos/cosmos-sdk/types"

// AjoloteToken represents a basic token on the Ajolocoin network.
// Each token has an ID, a characteristic of an axolotl and an amount of AJOL coins.
type AjoloteToken struct {
	ID             string // unique identifier of the token
	Characteristic string // axolotl characteristic associated with the token
	Amount         sdk.Coin
}

// AjolocoinApp is a minimal in-memory application that can mint tokens.
type AjolocoinApp struct {
	tokens map[string]AjoloteToken
}

// Setup creates the application with an empty token store.
func Setup() *AjolocoinApp {
	return &AjolocoinApp{
		tokens: make(map[string]AjoloteToken),
	}
}

// Mint creates a new AjoloteToken with the provided characteristic.
// The token is stored in memory and returned to the caller.
func (a *AjolocoinApp) Mint(id, characteristic string) AjoloteToken {
	token := AjoloteToken{
		ID:             id,
		Characteristic: characteristic,
		Amount:         sdk.NewInt64Coin("ajol", 1),
	}
	a.tokens[id] = token
	return token
}

// Get retrieves an AjoloteToken by its ID.
// It returns the token and a boolean indicating if it was found.
func (a *AjolocoinApp) Get(id string) (AjoloteToken, bool) {
	token, ok := a.tokens[id]
	return token, ok
}
