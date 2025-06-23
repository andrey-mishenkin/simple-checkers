package domain

import (
	"fmt"
	"math"
)

type Move struct {
	From Position
	To   Position
}

func NewMove(fromRow, fromCol, toRow, toCol int) Move {
	return Move{
		From: NewPosition(fromRow, fromCol),
		To:   NewPosition(toRow, toCol),
	}
}

func (m Move) String() string {
	return fmt.Sprintf("%d,%d:%d,%d", m.From.Row, m.From.Col, m.To.Row, m.To.Col)
}

type GameState int

const (
	GameInProgress GameState = iota
	GameWhiteWins
	GameBlackWins
	GameDraw
)

type Game struct {
	Board         *Board
	CurrentPlayer Player
	State         GameState
	MoveHistory   []Move
}

func NewGame() *Game {
	return &Game{
		Board:         NewBoard(),
		CurrentPlayer: PlayerWhite, // White always moves first
		State:         GameInProgress,
		MoveHistory:   make([]Move, 0),
	}
}

func (g *Game) IsValidMove(move Move) error {
	// basic
	if !move.From.IsValid() || !move.To.IsValid() {
		return fmt.Errorf("invalid positions")
	}

	piece := g.Board.GetPiece(move.From)
	if piece == nil {
		return fmt.Errorf("no piece at source position")
	}

	if piece.Player != g.CurrentPlayer {
		return fmt.Errorf("not your piece")
	}

	if !g.Board.IsEmpty(move.To) {
		return fmt.Errorf("destination is occupied")
	}

	// move is on dark squares only?
	if (move.From.Row+move.From.Col)%2 == 0 || (move.To.Row+move.To.Col)%2 == 0 {
		return fmt.Errorf("can only move on dark squares")
	}

	// move distance
	rowDiff := move.To.Row - move.From.Row
	colDiff := int(math.Abs(float64(move.To.Col - move.From.Col)))

	if colDiff != int(math.Abs(float64(rowDiff))) {
		return fmt.Errorf("must move diagonally")
	}

	// check direction constraint
	if !piece.IsKing() {
		if piece.Player == PlayerWhite && rowDiff > 0 {
			return fmt.Errorf("white pieces can only move up (towards row 0)")
		}
		if piece.Player == PlayerBlack && rowDiff < 0 {
			return fmt.Errorf("black pieces can only move down (towards row 7)")
		}
	}

	// simple move (1 square) or a jump (2 squares)
	if int(math.Abs(float64(rowDiff))) == 1 {
		// simple move
		return nil
	} else if int(math.Abs(float64(rowDiff))) == 2 {
		// jump - check if there's an opponent piece to capture
		midRow := move.From.Row + rowDiff/2
		midCol := move.From.Col + (move.To.Col-move.From.Col)/2
		midPos := NewPosition(midRow, midCol)

		midPiece := g.Board.GetPiece(midPos)
		if midPiece == nil {
			return fmt.Errorf("no piece to capture")
		}
		if midPiece.Player == piece.Player {
			return fmt.Errorf("cannot capture your own piece")
		}

		return nil
	}

	return fmt.Errorf("invalid move distance")
}

func (g *Game) MakeMove(move Move) error {
	if err := g.IsValidMove(move); err != nil {
		return fmt.Errorf("invalid move: %w", err)
	}

	piece := g.Board.GetPiece(move.From)

	// capturing move ?
	rowDiff := move.To.Row - move.From.Row
	if int(math.Abs(float64(rowDiff))) == 2 {
		// remove captured piece
		midRow := move.From.Row + rowDiff/2
		midCol := move.From.Col + (move.To.Col-move.From.Col)/2
		midPos := NewPosition(midRow, midCol)
		g.Board.RemovePiece(midPos)
	}

	g.Board.RemovePiece(move.From)
	g.Board.SetPiece(move.To, piece)

	// check for king
	// TODO:

	// Add move to history
	g.MoveHistory = append(g.MoveHistory, move)

	// Switch players
	g.CurrentPlayer = g.CurrentPlayer.Opponent()

	// Check for game end conditions
	g.updateGameState()

	return nil
}

// GetValidMoves returns all valid moves for the current player
func (g *Game) GetValidMoves() []Move {
	var moves []Move

	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			fromPos := NewPosition(row, col)
			piece := g.Board.GetPiece(fromPos)

			if piece == nil || piece.Player != g.CurrentPlayer {
				continue
			}

			// Check all possible diagonal moves
			directions := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

			for _, dir := range directions {
				// Check 1-square moves
				toPos := NewPosition(row+dir[0], col+dir[1])
				move := Move{From: fromPos, To: toPos}
				if g.IsValidMove(move) == nil {
					moves = append(moves, move)
				}

				// Check 2-square jumps
				toPos = NewPosition(row+dir[0]*2, col+dir[1]*2)
				move = Move{From: fromPos, To: toPos}
				if g.IsValidMove(move) == nil {
					moves = append(moves, move)
				}
			}
		}
	}

	return moves
}

func (g *Game) updateGameState() {
	whitePieces := 0
	blackPieces := 0

	// Count pieces
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			piece := g.Board.GetPiece(NewPosition(row, col))
			if piece != nil {
				if piece.Player == PlayerWhite {
					whitePieces++
				} else {
					blackPieces++
				}
			}
		}
	}

	// check win conditions
	if whitePieces == 0 {
		g.State = GameBlackWins
		return
	}
	if blackPieces == 0 {
		g.State = GameWhiteWins
		return
	}

	// current player has no valid moves?
	validMoves := g.GetValidMoves()
	if len(validMoves) == 0 {
		if g.CurrentPlayer == PlayerWhite {
			g.State = GameBlackWins
		} else {
			g.State = GameWhiteWins
		}
		return
	}

	g.State = GameInProgress
}

func (g *Game) IsGameOver() bool {
	return g.State != GameInProgress
}

func (g *Game) GetWinner() *Player {
	switch g.State {
	case GameWhiteWins:
		winner := PlayerWhite
		return &winner
	case GameBlackWins:
		winner := PlayerBlack
		return &winner
	default:
		return nil
	}
}
