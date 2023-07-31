package skullking

import (
	"reflect"
	"testing"
)

type TableTestWinner struct {
	Name  string
	Trick Trick

	Want int
}

func TestTrick_Winner(t *testing.T) {
	players := []*Player{
		{Name: "Victor"},
		{Name: "Erik"},
		{Name: "Ortega"},
		{Name: "Ignacio"},
	}

	table := []TableTestWinner{
		{
			Name: "All cards same suit @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 11}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "All cards same suit @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 9}},
				},
			},
			Want: 2,
		},
		{
			Name: "All cards same suit @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 9}},
				},
			},
			Want: 3,
		},
		{
			Name: "Highest Leading suit card wins @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 3,
		},
		{
			Name: "Highest Leading suit card wins @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 9}},
				},
			},
			Want: 2,
		},
		{
			Name: "One black card @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 9}},
				},
			},
			Want: 0,
		},
		{
			Name: "One black card @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 9}},
				},
			},
			Want: 1,
		},
		{
			Name: "One black card @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
				},
			},
			Want: 3,
		},
		{
			Name: "Highest black card wins @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 11}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 12}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "Highest black card wins @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 7}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 4}},
				},
			},
			Want: 1,
		},
		{
			Name: "Highest black card wins @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 4}},
				},
			},
			Want: 3,
		},
		{
			Name: "One escape @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeSuitGreen, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 4}},
				},
			},
			Want: 2,
		},
		{
			Name: "One escape @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitGreen, Value: 12}},
					{Player: players[1], Card: Card{Type: CardTypeSuitGreen, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeEscape}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "One escape @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitGreen, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeEscape}},
				},
			},
			Want: 1,
		},
		{
			Name: "Two escapes @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeEscape}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 4}},
				},
			},
			Want: 2,
		},
		{
			Name: "Two escapes @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeSuitGreen, Value: 10}},
					{Player: players[2], Card: Card{Type: CardTypeEscape}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 4}},
				},
			},
			Want: 3,
		},
		{
			Name: "Two escapes @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeEscape}},
				},
			},
			Want: 1,
		},
		{
			Name: "Three escapes @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeEscape}},
					{Player: players[2], Card: Card{Type: CardTypeEscape}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 4}},
				},
			},
			Want: 3,
		},
		{
			Name: "Three escapes @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeEscape}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeEscape}},
				},
			},
			Want: 2,
		},
		{
			Name: "All escapes @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeEscape}},
					{Player: players[2], Card: Card{Type: CardTypeEscape}},
					{Player: players[3], Card: Card{Type: CardTypeEscape}},
				},
			},
			Want: 0,
		},
		{
			Name: "Pirate wins suit @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "Pirate wins suit @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
				},
			},
			Want: 1,
		},
		{
			Name: "Pirate wins suit @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 3,
		},
		{
			Name: "Pirate wins black @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "Pirate wins black @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
				},
			},
			Want: 1,
		},
		{
			Name: "Pirate wins black @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 3,
		},
		{
			Name: "Two pirates",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 0,
		},
		{
			Name: "Mermaid wins suit @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "Mermaid wins suit @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[1], Card: Card{Type: CardTypeMermaid}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
				},
			},
			Want: 1,
		},
		{
			Name: "Mermaid wins suit @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 3,
		},
		{
			Name: "Mermaid wins black @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "Mermaid wins black @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[1], Card: Card{Type: CardTypeMermaid}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
				},
			},
			Want: 1,
		},
		{
			Name: "Mermaid wins black @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 3,
		},
		{
			Name: "Two mermaids",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 0,
		},
		{
			Name: "SkullKing wins suit @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "SkullKing wins suit @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[1], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
				},
			},
			Want: 1,
		},
		{
			Name: "SkullKing wins suit @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeSkullKing}},
				},
			},
			Want: 3,
		},
		{
			Name: "SkullKing wins black @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "SkullKing wins black @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[1], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
				},
			},
			Want: 1,
		},
		{
			Name: "SkullKing wins black @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeSkullKing}},
				},
			},
			Want: 3,
		},
		{
			Name: "Pirate beats mermaid @BEFORE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 0,
		},
		{
			Name: "Pirate beats mermaid @AFTER",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 3,
		},
		{
			Name: "SkullKing beats Pirate @BEFORE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 0,
		},
		{
			Name: "SkullKing beats Pirate @AFTER",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeSkullKing}},
				},
			},
			Want: 3,
		},
		{
			Name: "Mermaid beats SkullKing @BEFORE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeSkullKing}},
				},
			},
			Want: 0,
		},
		{
			Name: "Mermaid beats SkullKing @AFTER",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 3,
		},
		{
			Name: "Mermaid beats SkullKing AND Pirate @START",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypePirate}},
					{Player: players[3], Card: Card{Type: CardTypeSkullKing}},
				},
			},
			Want: 0,
		},
		{
			Name: "Mermaid beats SkullKing AND Pirate @MIDDLE",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypeMermaid}},
					{Player: players[3], Card: Card{Type: CardTypeSkullKing}},
				},
			},
			Want: 2,
		},
		{
			Name: "Mermaid beats SkullKing AND Pirate @END",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[2], Card: Card{Type: CardTypePirate}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 3,
		},
	}
	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.Trick.WinnerPosition()
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("Expected Player %d (%s) / Received Player %d (%s)", tt.Want, players[tt.Want].Name, got, players[got].Name)
			}
		})
	}

}

type TableTestPoints struct {
	Name  string
	Trick Trick

	Want int
}

func TestTrick_Points(t *testing.T) {
	players := []*Player{
		{Name: "Victor"},
		{Name: "Erik"},
		{Name: "Ortega"},
		{Name: "Ignacio"},
	}

	table := []TableTestWinner{
		{
			Name: "No points",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 11}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "One 14 card",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 9}},
				},
			},
			Want: 10,
		},
		{
			Name: "Two 14 cards",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 14}},
				},
			},
			Want: 20,
		},
		{
			Name: "14 Black card",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			Want: 20,
		},
		{
			Name: "14 Black and a 14 suit cards",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 9}},
				},
			},
			Want: 30,
		},
		{
			Name: "14 Black and two 14 suit cards",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 14}},
				},
			},
			Want: 40,
		},
		{
			Name: "One pirate so nothing",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 9}},
				},
			},
			Want: 0,
		},
		{
			Name: "Two pirates so nothing",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 0,
		},
		{
			Name: "One mermaid so nothing",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 9}},
				},
			},
			Want: 0,
		},
		{
			Name: "Two mermaids so nothing",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 0,
		},
		{
			Name: "One SkullKing so nothing",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 9}},
				},
			},
			Want: 0,
		},
		{
			Name: "Pirate and one Mermaid",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 12}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 20,
		},
		{
			Name: "Pirate and one Mermaid + 14s",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 30,
		},
		{
			Name: "Pirate and two Mermaid",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 40,
		},
		{
			Name: "Pirate and two Mermaid + 14",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 60,
		},
		{
			Name: "SkullKing and one Pirate",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 12}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeSkullKing}},
				},
			},
			Want: 30,
		},
		{
			Name: "SkullKing and one Pirate + 14s",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 60,
		},
		{
			Name: "SkullKing and two Pirate",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 60,
		},
		{
			Name: "SkullKing and two Pirate + 14",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 80,
		},
		{
			Name: "Mermaid and one SkullKing",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 12}},
					{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
					{Player: players[3], Card: Card{Type: CardTypeSkullKing}},
				},
			},
			Want: 40,
		},
		{
			Name: "Mermaid and one SkullKing + 14s",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 60,
		},
		{
			Name: "Mermaid SkullKing and Pirate",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 4}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 40,
		},
		{
			Name: "Mermaid SkullKing and Pirate + 14",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 50,
		},
		{
			Name: "Two Mermaids SkullKing and Pirate",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypeMermaid}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 40,
		},
		{
			Name: "Mermaid SkullKing and two Pirates",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypePirate}},
					{Player: players[3], Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 40,
		},
	}
	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.Trick.Points()
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("Expected Points %d / Received Points %d ", tt.Want, got)
			}
		})
	}

}

type TableTestBids struct {
	Name   string
	Round  Round
	Player Player

	Want int
}

func TestTrick_Bids(t *testing.T) {
	players := []*Player{
		{Name: "Victor"},
		{Name: "Erik"},
		{Name: "Ortega"},
		{Name: "Ignacio"},
	}

	table := []TableTestBids{
		{
			Name: "1st Round 0 Bid OK",
			Round: Round{
				Number: 1,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 0},
				},
			},
			Want: 10,
		},
		{
			Name: "1st Round 0 Bid KO",
			Round: Round{
				Number: 1,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 11}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 0},
				},
			},
			Want: -10,
		},
		{
			Name: "1st Round 1 Bid OK",
			Round: Round{
				Number: 1,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 13}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 1},
				},
			},
			Want: 20,
		},
		{
			Name: "1st Round 1 Bid KO",
			Round: Round{
				Number: 1,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitGreen, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 1},
				},
			},
			Want: -10,
		},
		{
			Name: "1st Round 1 Bid OK + Points",
			Round: Round{
				Number: 1,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 1},
				},
			},
			Want: 30,
		},
		{
			Name: "1st Round 1 Bid KO without Points",
			Round: Round{
				Number: 1,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 1},
				},
			},
			Want: -10,
		},
		{
			Name: "2nd Round 0 Bid OK",
			Round: Round{
				Number: 2,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
						},
					},
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 11}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 0},
				},
			},
			Want: 20,
		},
		{
			Name: "2nd Round 0 Bid KO won 1",
			Round: Round{
				Number: 2,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
						},
					},
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 11}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 0},
				},
			},
			Want: -20,
		},
		{
			Name: "2nd Round 0 Bid KO won 2",
			Round: Round{
				Number: 2,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
						},
					},
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSkullKing, Value: 11}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
							{Player: players[3], Card: Card{Type: CardTypePirate, Value: 4}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 0},
				},
			},
			Want: -20,
		},
		{
			Name: "2nd Round 1 Bid OK",
			Round: Round{
				Number: 2,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
						},
					},
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 11}},
							{Player: players[1], Card: Card{Type: CardTypeSuitGreen, Value: 14}},
							{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 1},
				},
			},
			Want: 20,
		},
		{
			Name: "2nd Round 1 Bid OK + points",
			Round: Round{
				Number: 2,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
						},
					},
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 11}},
							{Player: players[1], Card: Card{Type: CardTypeSuitGreen, Value: 14}},
							{Player: players[2], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 1},
				},
			},
			Want: 30,
		},
		{
			Name: "2nd Round 2 Bid OK",
			Round: Round{
				Number: 2,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
						},
					},
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 11}},
							{Player: players[1], Card: Card{Type: CardTypeSuitGreen, Value: 4}},
							{Player: players[2], Card: Card{Type: CardTypeEscape}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 2},
				},
			},
			Want: 40,
		},
		{
			Name: "2nd Round 2 Bid OK + points",
			Round: Round{
				Number: 2,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
						},
					},
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
							{Player: players[1], Card: Card{Type: CardTypePirate}},
							{Player: players[2], Card: Card{Type: CardTypePirate}},
							{Player: players[3], Card: Card{Type: CardTypePirate}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 2},
				},
			},
			Want: 140,
		},
		{
			Name: "2nd Round 1 Bid KO got 1",
			Round: Round{
				Number: 2,
				Tricks: []Trick{
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
							{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
							{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
							{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 14}},
						},
					},
					{
						Table: []*Play{
							{Player: players[0], Card: Card{Type: CardTypeMermaid}},
							{Player: players[1], Card: Card{Type: CardTypePirate}},
							{Player: players[2], Card: Card{Type: CardTypePirate}},
							{Player: players[3], Card: Card{Type: CardTypePirate}},
						},
					},
				},
				Bids: []Bid{
					{Player: *players[0], Bid: 2},
				},
			},
			Want: -10,
		},
	}
	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.Round.CheckBid(*players[0])
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("Expected Points %d / Received Points %d ", tt.Want, got)
			}
		})
	}

}

type TableTestLeading struct {
	name  string
	trick Trick
	want  CardType
}

func TestLeading(t *testing.T) {
	table := []TableTestLeading{
		{
			name: "single suit",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSuitGreen}},
					{Card: Card{Type: CardTypeSuitGreen}},
					{Card: Card{Type: CardTypeSuitGreen}},
				},
			},
			want: CardTypeSuitGreen,
		},
		{
			name: "three suit",
			trick: Trick{

				Table: []*Play{
					{Card: Card{Type: CardTypeSuitGreen}},
					{Card: Card{Type: CardTypeSuitYellow}},
					{Card: Card{Type: CardTypeSuitBlack}},
				},
			},
			want: CardTypeSuitGreen,
		},
		{
			name: "pirate",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypePirate}},
					{Card: Card{Type: CardTypeSuitYellow}},
					{Card: Card{Type: CardTypeSuitBlack}},
				},
			},
			want: CardTypeNone,
		},
		{
			name: "Escape",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeSuitYellow}},
					{Card: Card{Type: CardTypeSuitBlack}},
				},
			},
			want: CardTypeSuitYellow,
		},
		{
			name: "Escape two times",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeSuitBlack}},
				},
			},
			want: CardTypeSuitBlack,
		},
		{
			name: "Escape three times",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeEscape}},
				},
			},
			want: CardTypeNone,
		},
		{
			name: "Escape pirate suit",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypePirate}},
					{Card: Card{Type: CardTypeSuitGreen}},
				},
			},
			want: CardTypeNone,
		},
		{
			name: "Mermaid pirate suit",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeMermaid}},
					{Card: Card{Type: CardTypePirate}},
					{Card: Card{Type: CardTypeSuitGreen}},
				},
			},
			want: CardTypeNone,
		},
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.trick.Leading()
			if actual != tt.want {
				t.Fatalf("the suit is diferent,actual: \"%s\",expected:\"%s\" ", actual, tt.want)
			}
		})
	}
}
