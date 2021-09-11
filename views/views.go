package views


import "net/http"
import "fmt"
import "strconv"

import "github.com/gin-gonic/gin"

import . "api.card-deck/models"
import "api.card-deck/utils"


// store all card decks

var decks = []Deck {}


// view to create deck of cards
// card codes (AS, KH ...) can be passed (optional) thought the `cards` query to create specific cards in deck

func CreateDeck (c *gin.Context) {
	// get cards query parameter

	cards_query, query_exists := c.GetQuery("cards")

	// create instances for card suits

	var clubs, diamonds, hearts, spades []Card;

	// decide if query should be used or not
	// check if the `cards` parameter exists and if it has card codes (AS, KH ...)

	var use_query bool = query_exists && len(cards_query) > 0

	if (use_query) {
		// using `cards` to created decks

		clubs = CreateSuit("clubs", cards_query)
		diamonds = append(clubs, CreateSuit("diamonds", cards_query)...)
		hearts = append(diamonds, CreateSuit("hearts", cards_query)...)
		spades = append(hearts, CreateSuit("spades", cards_query)...)
	} else {
		// create deck with generic pattern of 52 cards

		clubs = CreateSuit("clubs")
		diamonds = append(clubs, CreateSuit("diamonds")...)
		hearts = append(diamonds, CreateSuit("hearts")...)
		spades = append(hearts, CreateSuit("spades")...)
	}

	// collect all card suits

	cards := spades

	// create instance of a new `Deck`
	// - `Deck_id` -> string -> unique UUID of each deck
	// - `Cards` -> [] Cards -> all cards available in deck
	// - `shuffled` -> bool -> if the deck has been shuffled or not
	// - `Remaining` -> int -> length of cards available in deck

	var new_deck Deck = Deck {
		Deck_id: utils.GenerateUniqueUUID(),
		Cards: cards,
		Shuffled: false,
		Remaining: len(cards),
	}

	// add `new_deck` to `decks`

	decks = append(decks, new_deck)

	// return `new_deck` as reponse to endpoint

	c.IndentedJSON(http.StatusOK, new_deck)
}


// view to fetch / open a specified deck using the `deck_id` -> UUID

func OpenDeck (c *gin.Context) {
	// get deck_id

	deck_id := c.Param("deck_id")

	// check if the deck_id length greater than zero, if not raise `StatusBadRequest` -> `bad_request`

	if (len(deck_id) <= 0) {
		c.IndentedJSON(http.StatusBadRequest, gin.H {"bad_request": "Deck ID is required"})
		return
	}

	// linear search for deck using `deck_id`

	for _, deck := range decks {
    if deck.Deck_id == deck_id {
        c.IndentedJSON(http.StatusOK, deck)
        return
    }
  }

  // raise `StatusNotFound` -> `not_found` if deck_id is not found

  c.IndentedJSON(http.StatusNotFound, gin.H {"not_found": "Deck not found"})
}


// draw cards from deck using;
// -`deck_id` (required)
// -`count` (required)

func DrawCard (c *gin.Context) {
	// get query parameters -> `deck_id` & `count`

	deck_id, deck_id_query_exists := c.GetQuery("deck_id")
	__count, count_query_exists := c.GetQuery("count")
	count, _ := strconv.Atoi(__count)

	// check if the required parameters are passed

	is_valid := deck_id_query_exists && count_query_exists

	// raise `StatusBadRequest` -> `bad_request` if required query parameters are not passed

	if (!is_valid) {
		c.IndentedJSON(http.StatusBadRequest, gin.H {"bad_request": "Deck ID and Card count is required"})
		return
	}

	// store cards with length of query parameter value of `count`

	var cards = []Card {}

	// linear search for deck using `deck_id`

	for _, deck := range decks {
    if deck.Deck_id == deck_id {

    	// get cards of length of query parameter value of `count`

    	for card_index := 0; card_index < count; card_index++ {

    		// update (card -> Card) to (cards []Card)

    		cards = append(cards, deck.Cards[card_index])
    	}

    	// return list of cards

    	c.IndentedJSON(http.StatusOK, cards)
    	return
    }
  }

  // raise `StatusNotFound` -> `not_found` if deck does not exist

  c.IndentedJSON(http.StatusNotFound, gin.H {"not_found": "Deck not found"})
}
