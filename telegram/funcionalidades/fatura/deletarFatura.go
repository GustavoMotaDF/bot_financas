package fatura

import (
	"fmt"
	"strconv"
	"strings"
	"telegram/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeleteFatura(bot *tgbotapi.BotAPI, userID *int64, chatID *int64, msg string, msgID *int) {
	idStr := strings.Split(msg, "_")
	valor := idStr[1]
	id, err := strconv.Atoi(valor)

	if err == nil {
		var id64 int64 = int64(id)
		err := repository.DeleteFatura(&id64)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(*chatID, "‼️ Fatura não localizada!"))
			fmt.Println(err)
			return
		}
		bot.Send(tgbotapi.NewEditMessageText(
			*chatID,
			*msgID,
			"✅ Fatura deletada com sucesso!",
		))
	}
}
