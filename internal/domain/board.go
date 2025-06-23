package domain

const BoardSize = 8

type Position struct {
	Row int
	Col int
}

func NewPosition(row, col int) Position {
	return Position{Row: row, Col: col}
}

func (p Position) IsValid() bool {
	return p.Row >= 0 && p.Row < BoardSize && p.Col >= 0 && p.Col < BoardSize
}

type Board struct {
	squares [BoardSize][BoardSize]*Piece
}

func NewBoard() *Board {
	board := &Board{}
	board.setupInitialPosition()
	return board
}

func (b *Board) setupInitialPosition() {
	// black pieces
	for row := 0; row < 3; row++ {
		for col := 0; col < BoardSize; col++ {
			if (row+col)%2 == 1 { // dark squares only
				b.squares[row][col] = NewPiece(PlayerBlack)
			}
		}
	}

	// white pieces
	for row := 5; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if (row+col)%2 == 1 { // dark squares only
				b.squares[row][col] = NewPiece(PlayerWhite)
			}
		}
	}
}

func (b *Board) GetPiece(pos Position) *Piece {
	if !pos.IsValid() {
		return nil
	}
	return b.squares[pos.Row][pos.Col]
}

func (b *Board) SetPiece(pos Position, piece *Piece) {
	if pos.IsValid() {
		b.squares[pos.Row][pos.Col] = piece
	}
}

func (b *Board) RemovePiece(pos Position) {
	if pos.IsValid() {
		b.squares[pos.Row][pos.Col] = nil
	}
}

func (b *Board) IsEmpty(pos Position) bool {
	return b.GetPiece(pos) == nil
}
