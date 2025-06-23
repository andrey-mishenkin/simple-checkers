package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/checkers/internal/domain"
)

type AIServiceInterface interface {
	GetBestMove(game *domain.Game) (domain.Move, error)
}

type AIService struct {
	random *rand.Rand
}

func NewAIService() AIServiceInterface {
	return &AIService{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GetBestMove returns the best move for the computer player, simple random strategy
func (ai *AIService) GetBestMove(game *domain.Game) (domain.Move, error) {
	validMoves := game.GetValidMoves()

	if len(validMoves) == 0 {
		return domain.Move{}, fmt.Errorf("no valid moves available")
	}

	// simple AI strategy: prioritize captures over regular moves
	captures := make([]domain.Move, 0)
	regularMoves := make([]domain.Move, 0)

	for _, move := range validMoves {
		// Check if it's a capture move (distance of 2)
		rowDiff := move.To.Row - move.From.Row
		if abs(rowDiff) == 2 {
			captures = append(captures, move)
		} else {
			regularMoves = append(regularMoves, move)
		}
	}

	// prefer captures if available
	if len(captures) > 0 {
		return captures[ai.random.Intn(len(captures))], nil
	}

	// make a random regular move
	return regularMoves[ai.random.Intn(len(regularMoves))], nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
