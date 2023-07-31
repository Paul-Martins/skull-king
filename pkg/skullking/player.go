package skullking

type Player struct {
	Name string
	Hand Cards
}

func (p *Player) cmp(player Player) bool {
	return p.Name == player.Name
}
