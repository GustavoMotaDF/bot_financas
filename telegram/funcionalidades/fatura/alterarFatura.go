package fatura

import (
	"fmt"
	"strings"
	"telegram/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func PagarFatura(bot *tgbotapi.BotAPI, userID *int64, chatID *int64, msg string, msgID *int) {
	idstr := strings.Split(msg, "_")
	id := ConvertParaInt(idstr[1])
	err := repository.PagarFatura(&id)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(*chatID, "‼️ Fatura não localizada!"))
		fmt.Println(err)
		return
	}
	bot.Send(tgbotapi.NewEditMessageText(
		*chatID,
		*msgID,
		fmt.Sprintf("✅ Fatura: %d, paga com sucesso!", id),
	))

}
