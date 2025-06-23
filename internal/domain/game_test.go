package domain

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	if game.CurrentPlayer != PlayerWhite {
		t.Errorf("Expected PlayerWhite to start, got %v", game.CurrentPlayer)
	}

	if game.State != GameInProgress {
		t.Errorf("Expected GameInProgress, got %v", game.State)
	}
}

func TestBoardInitialSetup(t *testing.T) {
	board := NewBoard()

	// black pieces are in correct positions
	blackCount := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < BoardSize; col++ {
			if (row+col)%2 == 1 { // Dark squares only
				piece := board.GetPiece(NewPosition(row, col))
				if piece == nil {
					t.Errorf("Expected black piece at (%d,%d), got nil", row, col)
				} else if piece.Player != PlayerBlack {
					t.Errorf("Expected black piece at (%d,%d), got %v", row, col, piece.Player)
				} else {
					blackCount++
				}
			}
		}
	}

	// white pieces are in correct positions
	whiteCount := 0
	for row := 5; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if (row+col)%2 == 1 { // Dark squares only
				piece := board.GetPiece(NewPosition(row, col))
				if piece == nil {
					t.Errorf("Expected white piece at (%d,%d), got nil", row, col)
				} else if piece.Player != PlayerWhite {
					t.Errorf("Expected white piece at (%d,%d), got %v", row, col, piece.Player)
				} else {
					whiteCount++
				}
			}
		}
	}

	// Should have 12 pieces each
	if blackCount != 12 {
		t.Errorf("Expected 12 black pieces, got %d", blackCount)
	}
	if whiteCount != 12 {
		t.Errorf("Expected 12 white pieces, got %d", whiteCount)
	}
}

func TestValidMove(t *testing.T) {
	game := NewGame()

	move := NewMove(5, 0, 4, 1) // Move white piece forward
	err := game.IsValidMove(move)
	if err != nil {
		t.Errorf("Expected valid move, got error: %v", err)
	}
}

func TestInvalidMove(t *testing.T) {
	game := NewGame()

	move := NewMove(3, 0, 4, 1) // Empty square
	err := game.IsValidMove(move)
	if err == nil {
		t.Error("Expected invalid move error, got nil")
	}
}

func TestMakeMove(t *testing.T) {
	game := NewGame()

	move := NewMove(5, 0, 4, 1)
	err := game.MakeMove(move)
	if err != nil {
		t.Errorf("Expected successful move, got error: %v", err)
	}

	// the piece moved?
	if !game.Board.IsEmpty(NewPosition(5, 0)) {
		t.Error("Source position should be empty after move")
	}

	piece := game.Board.GetPiece(NewPosition(4, 1))
	if piece == nil || piece.Player != PlayerWhite {
		t.Error("Destination should have white piece after move")
	}

	// turn switched
	if game.CurrentPlayer != PlayerBlack {
		t.Errorf("Expected turn to switch to PlayerBlack, got %v", game.CurrentPlayer)
	}
}
