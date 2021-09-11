package views

import "strconv"
import "strings"

import "api.card-deck/utils"


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


// method to create card suits
// - `cards_query` (options) can be passed to create specific cards in deck
// - `cards_query` is retrieved through positional argument `args` on index 0 / position 0 

func CreateSuit (card_type string, args ...string) []Card {
	// pseudo deck cards

	var __deck_cards = []string {"KING", "QUEEN", "JOKER"}

	// `lettered_cards` -> special cards needs to be specified, for slicing purposes

	var lettered_cards = []string {"KING", "QUEEN", "JOKER", "ACE"}

	// create 9 cards

	var card_ordinary_length int = 9

	// create array of 9 numbers (9 - 2)

	for i := card_ordinary_length; i >= 1; i-- {
		__deck_cards = append(__deck_cards, strconv.FormatInt(int64(i + 1), 10))
	}

	// add special card / `lettererd_card` -> "ACE" as last element

	__deck_cards = append(__deck_cards, "ACE")

	// real deck cards array

	var deck_cards = []Card {}

	// create cards from `__deck_cards`

	for card := 0; card < len(__deck_cards); card++ {
		_value := strings.ToUpper(__deck_cards[card])
		_suit := strings.ToUpper(card_type)
		_code := _value + _suit[0:1]

		// check if `_code` in `lettererd_cards`
		// slice first character of string if `_code` in `lettered_cards`

		if (utils.StringInSlice(lettered_cards, __deck_cards[card])) {
			_code = _value[0:1] + _suit[0:1]
		}

		// check if `cards_query` is in args

		if (len(args) > 0) {
			// - `cards_query` in args
			// using `cards_query` in args to create specified cards

			cards_query := strings.Split(args[0], ",")

			// to create specified cards in `cards_query`
			// check if `_code` in `cards_query`

			if (utils.StringInSlice(cards_query, _code)) {
				// update `deck_cards` array with `Card` instance

				deck_cards = append(
					deck_cards,
					Card {
						Value: _value,
						Suit: _suit,
						Code: _code,
					},
				)
			}
		} else {
			// - `cards_query` not in args
			// create default cards

			// update `deck_cards` array with `Card` instance

			deck_cards = append(
				deck_cards,
				Card {
					Value: _value,
					Suit: _suit,
					Code: _code,
				},
			)
		}
	}

	// return `deck_cards`

	return deck_cards
}
