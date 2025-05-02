package handlers

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ExtrairUserID(update *tgbotapi.Update) *int64 {

	if update.CallbackQuery != nil {
		return &update.CallbackQuery.From.ID
	}
	if update.Message != nil {
		return &update.Message.From.ID
	}
	return &Zero
}
func ExtrairChatID(update *tgbotapi.Update) *int64 {

	if update.CallbackQuery != nil {
		return &update.CallbackQuery.Message.Chat.ID
	}
	if update.Message != nil {
		return &update.Message.Chat.ID
	}
	return &Zero // ou -1, caso queira checar erros
}

func ExtrairMensagem(update *tgbotapi.Update) string {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Data
	}
	if update.Message != nil {
		return update.Message.Text
	} else {
		return ""
	}
}

func ExtrairMensagemID(update *tgbotapi.Update) *int {
	if update.CallbackQuery != nil {
		return &update.CallbackQuery.Message.MessageID
	}
	if update.Message != nil {
		return &update.Message.MessageID
	} else {
		return nil
	}
}
func ConvertParaBool(valor string) (bool, error) {
	val, err := strconv.ParseBool(valor)
	if err != nil {
		fmt.Println("Erro ao converter:", err)
		return false, err
	}
	return val, nil
}
func ConvertParaInt(valor string) int {
	mes, err := strconv.Atoi(valor)
	if err != nil {
		fmt.Println("Erro ao converter em string")
		return -1
	}
	return mes
}
