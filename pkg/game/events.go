package game

import (
	"github.com/metalblueberry/skull-king/pkg/skullking"
)

type EventCardPlayed struct {
	Player skullking.Player
	Card   skullking.Card
}

func (e EventCardPlayed) SaySomething() {

}

type Event interface {
	SaySomething()
}

func (e EventCardPlayed) Send(context Context) {
	Subscribe(context, func(Event EventCardPlayed) {

	})
}
