package domain

type PieceType int

const (
	Regular PieceType = iota
	King              // stubbed - not implemented in game logic
)

type Piece struct {
	Player Player
	Type   PieceType
}

func NewPiece(player Player) *Piece {
	return &Piece{
		Player: player,
		Type:   Regular,
	}
}

func (p *Piece) IsKing() bool {
	return p.Type == King
}

func (p *Piece) Symbol() string {
	return p.Player.Symbol()
}
