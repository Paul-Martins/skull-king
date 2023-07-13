package skullking

// Trick Consists of each player, in clockwise order, playing
// 1 card face-up to the table. The person who plays the
// highest valued card wins and takes the trick. The player
// gathers the cards in a stack before them.
type Trick struct {
	Table []*Play
}

// Round Consisting of 1 or more tricks. The number of
// tricks in a round is determined by the number of cards
// dealt. A round begins by dealing cards and ends when all
// cards dealt have been played.
type Round struct {
	Tricks []Trick
}

// Play is the Card played by a Player during a Trick
type Play struct {
	Player *Player
	Card   Card
}

// NewTrick creates a new empty Trick struct that can be used to Play cards by the Players
func NewTrick(numberOfPlayers int) *Trick {
	return &Trick{
		Table: make([]*Play, numberOfPlayers),
	}
}

// Play adds the Play action for the next player to the Trick
// It will return an error if the card cannot be played
func (t *Trick) Play(p Play) error {
	panic("Not Implemented Yet")
}

// Winner will return the player that wins the current Trick
func (t *Trick) Winner() *Player {

	//t.Table[0].
	return nil
}

// Points returns the amount of points that this specific trick is worth for the player that wins it.
func (t *Trick) Points() int {
	return 0
}

func (t *Trick) Leading() CardType {
	if t.Table[0].Card.Type != CardTypeSuitGreen &&
		t.Table[0].Card.Type != CardTypeSuitBlack &&
		t.Table[0].Card.Type != CardTypeSuitYellow &&
		t.Table[0].Card.Type != CardTypeSuitPurple {
		return CardTypeNone
	}
	return t.Table[0].Card.Type
}