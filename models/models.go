package models

// `Card` object / model

type Card struct {
	Value string 	`json:"value"`
	Suit	string 	`json:"suit"`
	Code  string 	`json:"code"`
}

// `ProxyDeck` object / model

type ProxyDeck struct {
	Cards 		[]Card 	`json:"cards"`
}

// `Deck` object / model

type Deck struct {
	Deck_id		string 	`json:"deck_id"`
	Shuffled	bool 		`json:"shuffled"`
	Remaining int 		`json:"remaining"`
	Cards 		[]Card 	`json:"cards,omitempty"`
}
