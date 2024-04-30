package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card represents a single playing card with a suit and a value.
type Card struct {
	Suit  string
	Value string
}

// Deck represents a deck of cards.
type Deck []Card

// NewDeck initializes and returns a new deck of cards.
func NewDeck() Deck {
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	deck := make(Deck, 0, 52)
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}
	return deck
}

// Shuffle randomizes the order of cards in the deck.
func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

// Draw removes the top card from the deck and returns it.
func (d *Deck) Draw() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

// Player represents a blackjack player or the dealer.
type Player struct {
	Hand []Card
}

// AddCard adds a card to the player's hand.
func (p *Player) AddCard(card Card) {
	p.Hand = append(p.Hand, card)
}

// Score calculates and returns the player's score.
func (p *Player) Score() int {
	score := 0
	aces := 0
	for _, card := range p.Hand {
		switch card.Value {
		case "Ace":
			aces++
			score += 11
		case "King", "Queen", "Jack", "10":
			score += 10
		default:
			// Convert card value to integer for 2-9
			var val int
			fmt.Sscanf(card.Value, "%d", &val)
			score += val
		}
	}
	// Adjust score for Aces if needed
	for aces > 0 && score > 21 {
		score -= 10
		aces--
	}
	return score
}

func main() {
	// Create a deck and shuffle it
	deck := NewDeck()
	deck.Shuffle()

	// Create player and dealer
	player := Player{}
	dealer := Player{}

	// Initial two cards draw for player and dealer
	player.AddCard(deck.Draw())
	player.AddCard(deck.Draw())
	dealer.AddCard(deck.Draw())
	dealer.AddCard(deck.Draw())

	// Display initial hands (in a real game, dealer's one card might be hidden)
	fmt.Println("Player's Hand:", player.Hand)
	fmt.Println("Dealer's Hand:", dealer.Hand)
}
