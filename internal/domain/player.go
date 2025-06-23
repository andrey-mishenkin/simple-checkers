package domain

type Player int

const (
	PlayerWhite Player = iota
	PlayerBlack
)

func (p Player) String() string {
	switch p {
	case PlayerWhite:
		return "White (human)"
	case PlayerBlack:
		return "Black (AI)"
	default:
		return "Unknown"
	}
}

func (p Player) Symbol() string {
	switch p {
	case PlayerWhite:
		return "W"
	case PlayerBlack:
		return "B"
	default:
		return "?"
	}
}

func (p Player) Opponent() Player {
	switch p {
	case PlayerWhite:
		return PlayerBlack
	case PlayerBlack:
		return PlayerWhite
	default:
		return PlayerWhite
	}
}
