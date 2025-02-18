package skullking

// Trick Consists of each player, in clockwise order, playing
// 1 card face-up to the table. The person who plays the
// highest valued card wins and takes the trick. The player
// gathers the cards in a stack before them.
type Trick struct {
	Table []*Play
}

//Structure to gather the necessary data to declare a winner
type InfoWinner struct {
	SkullKing  int
	Pirate     int
	Mermaid    int
	Black      int
	BlackValue int
	Suit       int
	SuitValue  int
}

//Structure to gather the necessary data to count how many points this trick would concede
type InfoPoints struct {
	SkullKing bool
	Pirates   int
	Mermaids  int
	Black14   bool
	Suits14   int
}

// Round Consisting of 1 or more tricks. The number of
// tricks in a round is determined by the number of cards
// dealt. A round begins by dealing cards and ends when all
// cards dealt have been played.
type Round struct {
	Number int
	Tricks []Trick
	Bids   []Bid
}

// Play is the Card played by a Player during a Trick
type Play struct {
	Player *Player
	Card   Card
}

type Bid struct {
	Player Player
	Bid    int
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

func (r *Round) CheckBid(player Player) int {
	bid := r.getBidByPlayer(player)
	roundsWon := 0
	pointsWon := 0
	for _, t := range r.Tricks {
		if t.Winner().cmp(player) {
			roundsWon++
			pointsWon += t.Points()
		}
	}
	if bid == 0 && roundsWon > 0 {
		return -r.Number * 10
	}
	if bid == 0 {
		return r.Number * 10
	}
	if bid == roundsWon {
		pointsWon += bid * 20
		return pointsWon
	}
	difference := bid - roundsWon
	if difference > 0 {
		difference = -difference
	}
	return difference * 10

}

func (r *Round) getBidByPlayer(player Player) int {
	for _, b := range r.Bids {
		if player.cmp(b.Player) {
			return b.Bid
		}
	}
	return 0
}

//Loop over the deck to gather info to check points
func (t *Trick) GatheringInfoPoints() InfoPoints {
	info := InfoPoints{
		SkullKing: false,
		Pirates:   0,
		Mermaids:  0,
		Black14:   false,
		Suits14:   0,
	}
	for _, v := range t.Table {
		if v.Card.Type == CardTypeSkullKing {
			info.SkullKing = true
			continue
		}
		if v.Card.Type == CardTypePirate {
			info.Pirates++
			continue
		}
		if v.Card.Type == CardTypeMermaid {
			info.Mermaids++
			continue
		}
		if v.Card.Value == 14 {
			if v.Card.Type == CardTypeSuitBlack {
				info.Black14 = true
			} else {
				info.Suits14++
			}
		}
	}
	return info
}

// Points returns the amount of points that this specific trick is worth for the player that wins it.
func (t *Trick) Points() int {
	info := t.GatheringInfoPoints()
	sum := 0
	sum += info.Suits14 * 10
	if info.Black14 {
		sum += 20
	}
	if info.SkullKing && info.Mermaids > 0 {
		sum += 40
		return sum
	}
	if info.SkullKing && info.Pirates > 0 {
		sum += info.Pirates * 30
		return sum
	}
	if info.Pirates > 0 && info.Mermaids > 0 {
		sum += info.Mermaids * 20
		return sum
	}

	return sum
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

//Loop over the deck to gather info to check winner
func (t *Trick) GatheringInfoWinner() InfoWinner {
	info := InfoWinner{
		SkullKing:  -1,
		Pirate:     -1,
		Mermaid:    -1,
		Black:      -1,
		BlackValue: 0,
		Suit:       -1,
		SuitValue:  0,
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

// Winner will return the player that wins the current Trick
func (t *Trick) WinnerPosition() int {
	info := t.GatheringInfoWinner()
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

func (t *Trick) Winner() *Player {
	return t.Table[t.WinnerPosition()].Player
}
