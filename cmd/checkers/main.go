package main

import (
	"fmt"
	"log"

	"github.com/checkers/internal/handler"
	"github.com/checkers/internal/service"
)

func main() {
	// init services
	gameService := service.NewGameService()
	aiService := service.NewAIService()

	// init handler
	gameHandler := handler.NewGameHandler(gameService, aiService)

	if err := gameHandler.RunGame(); err != nil {
		log.Printf("Game ended with error: %v", err)
		return
	}

	fmt.Println("Game completed successfully!")
}
