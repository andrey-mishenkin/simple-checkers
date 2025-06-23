package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/checkers/internal/domain"
	"github.com/checkers/internal/ui"
)

// GameServiceInterface defines the contract for game operations
type GameServiceInterface interface {
	NewGame() *domain.Game
	ParseMove(moveStr string) (domain.Move, error)
	MakeMove(game *domain.Game, move domain.Move) error
	GetBoard(game *domain.Game) string
	GetValidMoves(game *domain.Game) []domain.Move
	IsGameOver(game *domain.Game) bool
	GetCurrentPlayer(game *domain.Game) domain.Player
	GetGameStatus(game *domain.Game) string
}

// GameService implements GameServiceInterface
type GameService struct {
	boardRenderer *ui.BoardRenderer
}

func NewGameService() GameServiceInterface {
	return &GameService{
		boardRenderer: ui.NewBoardRenderer(),
	}
}

func (s *GameService) NewGame() *domain.Game {
	return domain.NewGame()
}

func (s *GameService) ParseMove(moveStr string) (domain.Move, error) {
	moveStr = strings.TrimSpace(moveStr)

	parts := strings.Split(moveStr, ":")
	if len(parts) != 2 {
		return domain.Move{}, fmt.Errorf("invalid move format, expected 'fromRow,fromCol:toRow,toCol'")
	}

	fromParts := strings.Split(parts[0], ",")
	toParts := strings.Split(parts[1], ",")

	if len(fromParts) != 2 || len(toParts) != 2 {
		return domain.Move{}, fmt.Errorf("invalid move format, expected 'fromRow,fromCol:toRow,toCol'")
	}

	fromRow, err := strconv.Atoi(strings.TrimSpace(fromParts[0]))
	if err != nil {
		return domain.Move{}, fmt.Errorf("invalid from row: %w", err)
	}

	fromCol, err := strconv.Atoi(strings.TrimSpace(fromParts[1]))
	if err != nil {
		return domain.Move{}, fmt.Errorf("invalid from column: %w", err)
	}

	toRow, err := strconv.Atoi(strings.TrimSpace(toParts[0]))
	if err != nil {
		return domain.Move{}, fmt.Errorf("invalid to row: %w", err)
	}

	toCol, err := strconv.Atoi(strings.TrimSpace(toParts[1]))
	if err != nil {
		return domain.Move{}, fmt.Errorf("invalid to column: %w", err)
	}

	return domain.NewMove(fromRow, fromCol, toRow, toCol), nil
}

func (s *GameService) MakeMove(game *domain.Game, move domain.Move) error {
	return game.MakeMove(move)
}

func (s *GameService) GetBoard(game *domain.Game) string {
	return s.boardRenderer.RenderBoard(game.Board)
}

func (s *GameService) GetValidMoves(game *domain.Game) []domain.Move {
	return game.GetValidMoves()
}

func (s *GameService) IsGameOver(game *domain.Game) bool {
	return game.IsGameOver()
}

func (s *GameService) GetCurrentPlayer(game *domain.Game) domain.Player {
	return game.CurrentPlayer
}

func (s *GameService) GetGameStatus(game *domain.Game) string {
	if game.IsGameOver() {
		if winner := game.GetWinner(); winner != nil {
			return fmt.Sprintf("Game Over! %s wins!", winner.String())
		}
		return "Game Over! It's a draw!"
	}

	return fmt.Sprintf("Current player: %s", game.CurrentPlayer.String())
}
