package handler

import (
	"fmt"
	"strings"

	"github.com/checkers/internal/domain"
	"github.com/checkers/internal/service"
	"github.com/checkers/internal/ui"
)

type GameHandler struct {
	gameService service.GameServiceInterface
	aiService   service.AIServiceInterface
}

func NewGameHandler(gameService service.GameServiceInterface, aiService service.AIServiceInterface) *GameHandler {
	return &GameHandler{
		gameService: gameService,
		aiService:   aiService,
	}
}

func (h *GameHandler) RunGame() error {
	game := h.gameService.NewGame()

	ui.PrintTitle()
	ui.PrintWelcome()

	for !h.gameService.IsGameOver(game) {

		// show board
		fmt.Println(h.gameService.GetBoard(game))
		fmt.Println(h.gameService.GetGameStatus(game))

		currentPlayer := h.gameService.GetCurrentPlayer(game)

		if currentPlayer == domain.PlayerWhite {
			// human turn
			if err := h.handleHumanTurn(game); err != nil {
				return fmt.Errorf("error handling human turn: %w", err)
			}
		} else {
			// computer turn
			if err := h.handleComputerTurn(game); err != nil {
				return fmt.Errorf("error handling computer turn: %w", err)
			}
		}
	}

	// game over
	fmt.Println(h.gameService.GetBoard(game))
	fmt.Println(h.gameService.GetGameStatus(game))
	fmt.Println("Game over!")

	return nil
}

func (h *GameHandler) handleHumanTurn(game *domain.Game) error {
	for {
		fmt.Print("\nYour move (White): ")
		input, err := ui.ReadInput()
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.TrimSpace(input)

		switch strings.ToLower(input) {
		case "quit", "exit", "q":
			return fmt.Errorf("game quit by user")
		case "help", "h":
			ui.PrintHelp()
			continue
		case "":
			fmt.Println("Please enter a move.")
			continue
		}

		// parse and make the move
		move, err := h.gameService.ParseMove(input)
		if err != nil {
			fmt.Printf("Invalid move format: %v\n", err)
			continue
		}

		if err := h.gameService.MakeMove(game, move); err != nil {
			fmt.Printf("Invalid move: %v\n", err)
			continue
		}

		fmt.Printf("You moved: %s\n", move.String())
		break
	}

	return nil
}

func (h *GameHandler) handleComputerTurn(game *domain.Game) error {
	fmt.Println("\nComputer is thinking...")

	move, err := h.aiService.GetBestMove(game)
	if err != nil {
		return fmt.Errorf("AI failed to find a move: %w", err)
	}

	if err := h.gameService.MakeMove(game, move); err != nil {
		return fmt.Errorf("AI made invalid move: %w", err)
	}

	fmt.Printf("Computer moved: %s\n", move.String())
	return nil
}
