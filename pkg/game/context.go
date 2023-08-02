package game

import "github.com/metalblueberry/skull-king/pkg/skullking"

type Context struct {
	Players []*skullking.Player
	Deck    skullking.Deck
	Rounds  []*skullking.Round
}

func Subscribe[T Event](context Context, handler func(T)) {

}
