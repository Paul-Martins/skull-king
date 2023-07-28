package skullking

// Trick Consists of each player, in clockwise order, playing
// 1 card face-up to the table. The person who plays the
// highest valued card wins and takes the trick. The player
// gathers the cards in a stack before them.
type Trick struct {
	Table []*Play
}
type Info struct {
	SkullKing  int
	Pirate     int
	Mermaid    int
	Black      int
	BlackValue int
	Suit       int
	SuitValue  int
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
	winnerPosition := t.Winner2()
	return t.Table[winnerPosition].Player
}

// Points returns the amount of points that this specific trick is worth for the player that wins it.
func (t *Trick) Points() int {
	return 0
}

func (t *Trick) Leading() CardType {

	position := 0
	for i := position; i < len(t.Table); i++ {
		if t.Table[i].Card.Type != CardTypeEscape {
			position = i
			break
		}
	}

	if t.Table[position].Card.Type != CardTypeSuitGreen &&
		t.Table[position].Card.Type != CardTypeSuitBlack &&
		t.Table[position].Card.Type != CardTypeSuitYellow &&
		t.Table[position].Card.Type != CardTypeSuitPurple {
		return CardTypeNone
	}

	return t.Table[position].Card.Type
}

func (t *Trick) GatheringInfo() Info {
	info := Info{
		SkullKing:  -1,
		Pirate:     -1,
		Mermaid:    -1,
		Black:      -1,
		BlackValue: -1,
		Suit:       -1,
		SuitValue:  -1,
	}
	leadingSuit := t.Leading()
	for i, v := range t.Table {
		if v.Card.Type == CardTypeSkullKing {
			info.SkullKing = i
			continue
		}
		if info.Pirate == -1 && v.Card.Type == CardTypePirate {
			info.Pirate = i
			continue
		}
		if info.Mermaid == -1 && v.Card.Type == CardTypeMermaid {
			info.Mermaid = i
			continue
		}
		if v.Card.Type == CardTypeSuitBlack && v.Card.Value > info.BlackValue {
			info.Black = i
			info.BlackValue = v.Card.Value
			continue
		}
		if v.Card.Type == leadingSuit && v.Card.Value > info.SuitValue {
			info.Suit = i
			info.SuitValue = v.Card.Value
			continue
		}
	}
	return info
}

func (t *Trick) Winner2() int {
	info := t.GatheringInfo()
	if info.SkullKing >= 0 && info.Mermaid >= 0 {
		return info.Mermaid
	}
	if info.SkullKing >= 0 && info.Pirate >= 0 {
		return info.SkullKing
	}
	if info.Pirate >= 0 && info.Mermaid >= 0 {
		return info.Pirate
	}
	if info.SkullKing >= 0 {
		return info.SkullKing
	}
	if info.Pirate >= 0 {
		return info.Pirate
	}
	if info.Mermaid >= 0 {
		return info.Mermaid
	}
	if info.Black >= 0 {
		return info.Black
	}
	if info.Suit >= 0 {
		return info.Suit
	}

	return 0
}
