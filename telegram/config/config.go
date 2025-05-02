package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type TelegramConfig struct {
	ApiKey string
	UserID []int64 // Agora é uma lista de int64
}

var AppConfig TelegramConfig

func LoadConfig() {
	apiKey := os.Getenv("API_KEY")
	userIDStr := os.Getenv("TELEGRAM_USER_IDS") // Variável de ambiente que passa a lista de IDs

	if apiKey == "" || userIDStr == "" {
		log.Fatal("API_KEY ou TELEGRAM_USER_IDS não definidos")
	}

	// Divida a string com IDs separados por vírgula
	userIDs := strings.Split(userIDStr, ",")

	// Converta cada item para int64 e adicione à slice
	var userIDList []int64
	for _, idStr := range userIDs {
		userID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Fatalf("Erro ao converter um UserID: %v", err)
		}
		userIDList = append(userIDList, userID)
	}

	// Atribuindo valores ao AppConfig
	AppConfig = TelegramConfig{
		ApiKey: apiKey,
		UserID: userIDList, // Lista de IDs
	}
}
