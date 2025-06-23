package ui

import (
	"fmt"
	"strings"

	"github.com/checkers/internal/domain"
)

// BoardRenderer handles the visual representation of the game board
type BoardRenderer struct{}

// NewBoardRenderer creates a new board renderer
func NewBoardRenderer() *BoardRenderer {
	return &BoardRenderer{}
}

// RenderBoard returns a string representation of the board
func (r *BoardRenderer) RenderBoard(board *domain.Board) string {
	var sb strings.Builder

	sb.WriteString("  0 1 2 3 4 5 6 7\n")
	for row := 0; row < domain.BoardSize; row++ {
		sb.WriteString(fmt.Sprintf("%d ", row))
		for col := 0; col < domain.BoardSize; col++ {
			piece := board.GetPiece(domain.NewPosition(row, col))
			if piece != nil {
				sb.WriteString(piece.Symbol())
			} else if (row+col)%2 == 1 {
				sb.WriteString(".")
			} else {
				sb.WriteString(" ")
			}
			sb.WriteString(" ")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
