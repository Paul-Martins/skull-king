package skullking

import "math/rand"

type Deck struct {
	Cards []Card
}

type Card struct {
	// Number identifies the card uniquely with a number
	Number int
	// Type of the card
	Type CardType
	// Value of the card if needed
	Value int
}

// que es un mazo un frenesi una sombra una ilusion

// cartas para todos de typo carta
// Los comentarios son terrorismo AT:Victor
// calla ya pesado
type Cards []Card

type CardValue int
type CardType string

const (
	CardTypeNone CardType = ""

	CardTypeSkullKing CardType = "SkullKing"
	CardTypeMermaid   CardType = "Mermaid"
	CardTypePirate    CardType = "Pirate"
	CardTypeEscape    CardType = "Escape"
	CardTypeTigress   CardType = "Tigress"

	CardTypeSuitGreen  CardType = "green"
	CardTypeSuitYellow CardType = "yellow"
	CardTypeSuitPurple CardType = "purple"
	CardTypeSuitBlack  CardType = "black"
)

var (
	OriginalDeck = Deck{
		Cards: []Card{
			{Number: 1, Type: CardTypeSuitGreen, Value: 1},
			{Number: 2, Type: CardTypeSuitGreen, Value: 2},
			{Number: 3, Type: CardTypeSuitGreen, Value: 3},
			{Number: 4, Type: CardTypeSuitGreen, Value: 4},
			{Number: 5, Type: CardTypeSuitGreen, Value: 5},
			{Number: 6, Type: CardTypeSuitGreen, Value: 6},
			{Number: 7, Type: CardTypeSuitGreen, Value: 7},
			{Number: 8, Type: CardTypeSuitGreen, Value: 8},
			{Number: 9, Type: CardTypeSuitGreen, Value: 9},
			{Number: 10, Type: CardTypeSuitGreen, Value: 10},
			{Number: 11, Type: CardTypeSuitGreen, Value: 11},
			{Number: 12, Type: CardTypeSuitGreen, Value: 12},
			{Number: 13, Type: CardTypeSuitGreen, Value: 13},
			{Number: 14, Type: CardTypeSuitGreen, Value: 14},
			{Number: 15, Type: CardTypeSuitYellow, Value: 1},
			{Number: 16, Type: CardTypeSuitYellow, Value: 2},
			{Number: 17, Type: CardTypeSuitYellow, Value: 3},
			{Number: 18, Type: CardTypeSuitYellow, Value: 4},
			{Number: 19, Type: CardTypeSuitYellow, Value: 5},
			{Number: 20, Type: CardTypeSuitYellow, Value: 6},
			{Number: 21, Type: CardTypeSuitYellow, Value: 7},
			{Number: 22, Type: CardTypeSuitYellow, Value: 8},
			{Number: 23, Type: CardTypeSuitYellow, Value: 9},
			{Number: 24, Type: CardTypeSuitYellow, Value: 10},
			{Number: 25, Type: CardTypeSuitYellow, Value: 11},
			{Number: 26, Type: CardTypeSuitYellow, Value: 12},
			{Number: 27, Type: CardTypeSuitYellow, Value: 13},
			{Number: 28, Type: CardTypeSuitYellow, Value: 14},
			{Number: 29, Type: CardTypeSuitBlack, Value: 1},
			{Number: 30, Type: CardTypeSuitBlack, Value: 2},
			{Number: 31, Type: CardTypeSuitBlack, Value: 3},
			{Number: 32, Type: CardTypeSuitBlack, Value: 4},
			{Number: 33, Type: CardTypeSuitBlack, Value: 5},
			{Number: 34, Type: CardTypeSuitBlack, Value: 6},
			{Number: 35, Type: CardTypeSuitBlack, Value: 7},
			{Number: 36, Type: CardTypeSuitBlack, Value: 8},
			{Number: 37, Type: CardTypeSuitBlack, Value: 9},
			{Number: 38, Type: CardTypeSuitBlack, Value: 10},
			{Number: 39, Type: CardTypeSuitBlack, Value: 11},
			{Number: 40, Type: CardTypeSuitBlack, Value: 12},
			{Number: 41, Type: CardTypeSuitBlack, Value: 13},
			{Number: 42, Type: CardTypeSuitBlack, Value: 14},
			{Number: 43, Type: CardTypeSuitPurple, Value: 1},
			{Number: 44, Type: CardTypeSuitPurple, Value: 2},
			{Number: 45, Type: CardTypeSuitPurple, Value: 3},
			{Number: 46, Type: CardTypeSuitPurple, Value: 4},
			{Number: 47, Type: CardTypeSuitPurple, Value: 5},
			{Number: 48, Type: CardTypeSuitPurple, Value: 6},
			{Number: 49, Type: CardTypeSuitPurple, Value: 7},
			{Number: 50, Type: CardTypeSuitPurple, Value: 8},
			{Number: 51, Type: CardTypeSuitPurple, Value: 9},
			{Number: 52, Type: CardTypeSuitPurple, Value: 10},
			{Number: 53, Type: CardTypeSuitPurple, Value: 11},
			{Number: 54, Type: CardTypeSuitPurple, Value: 12},
			{Number: 55, Type: CardTypeSuitPurple, Value: 13},
			{Number: 56, Type: CardTypeSuitPurple, Value: 14},

			{Number: 57, Type: CardTypePirate},
			{Number: 58, Type: CardTypePirate},
			{Number: 59, Type: CardTypePirate},
			{Number: 60, Type: CardTypePirate},
			{Number: 61, Type: CardTypePirate},
			{Number: 62, Type: CardTypeTigress},
			{Number: 63, Type: CardTypeSkullKing},
			{Number: 64, Type: CardTypeMermaid},
			{Number: 65, Type: CardTypeMermaid},
			{Number: 66, Type: CardTypeEscape},
			{Number: 67, Type: CardTypeEscape},
			{Number: 68, Type: CardTypeEscape},
			{Number: 69, Type: CardTypeEscape},
			{Number: 70, Type: CardTypeEscape},
		},
	}
)

func NewDeck() Deck {
	return Deck{
		Cards: make([]Card, 0),
	}
}

func (d *Deck) Put(cards ...Card) {
	d.Cards = append(d.Cards, cards...)
}

func (d *Deck) Draw(n int) []Card {

	//d.Cards[len(d.Cards)-1]
	hand := d.Cards[len(d.Cards)-2:] //len(d.Cards)
	deck2 := d.Cards[:len(d.Cards)-2]
	d.Cards = deck2
	return hand
}

func (d *Deck) Shufle() {
	size := len(d.Cards)

	for i := 0; i < 7322; i++ {
		n1 := rand.Intn(size)
		n2 := rand.Intn(size)
		cardtem := d.Cards[n1]
		cardtem2 := d.Cards[n2]

		d.Cards[n1] = cardtem2
		d.Cards[n2] = cardtem
	}

}

func (d *Deck) Deal(hands int, cards int) []Cards {
	return nil
}
