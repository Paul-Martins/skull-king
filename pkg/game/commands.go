package game

import "github.com/metalblueberry/skull-king/pkg/skullking"

type CommandPlayCard struct {
	Player skullking.Player
	Card   skullking.Card
}

type CommandBid struct {
	Player skullking.Player
	Amount int
}

type CommandType string

const (
	CommandTypePlayCard = "PlayCard"
	CommandTypeBid      = "Bid"
)

func (c CommandPlayCard) CommandType() CommandType {
	return CommandTypePlayCard
}

func (c CommandBid) CommandType() CommandType {
	return CommandTypeBid
}

func (c CommandPlayCard) Execute(context *Context) {
}

func (c CommandBid) Execute(context *Context) {
}

type Command interface {
	CommandType() CommandType
	Execute(context *Context)
}

func SendCommand(command Command) error {
	panic("not yet")
}
