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
			got := tt.Trick.Winner()
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("Expected Player %d (%s) / Received Player %d (%s)", tt.Want, players[tt.Want].Name, got, players[got].Name)
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
