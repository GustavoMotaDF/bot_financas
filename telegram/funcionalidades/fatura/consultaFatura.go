package fatura

import (
	"fmt"
	"strings"
	"telegram/models"
	"telegram/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetAllFaturas(bot *tgbotapi.BotAPI, chatID *int64, pagas bool) {
	faturas, err := repository.GetAllFaturas(pagas)
	if len(faturas) == 0 || err != nil {
		bot.Send(tgbotapi.NewMessage(*chatID, "‼️ Sem faturas para mostrar"))
		fmt.Println(err)
		return
	}
	bot.Send(tgbotapi.NewMessage(*chatID, "Foram localizadas essas faturas: "))
	for _, item := range faturas {
		var id = int64(item.ID)
		BotoesFatura(bot, chatID, id, models.ModelaFatura(&item), pagas)
	}
}

func GetFaturasMes(bot *tgbotapi.BotAPI, chatID *int64, msg string, paga bool) {
	///consultar_faturas_mes_4_2025_paga_true // false
	mesAno := strings.Split(msg, "_")
	var mes = ConvertParaInt(mesAno[3])
	var ano = ConvertParaInt(mesAno[4])

	faturas, err := repository.GetFaturasMes(mes, ano, paga)
	if len(faturas) == 0 || err != nil {
		bot.Send(tgbotapi.NewMessage(*chatID, "‼️ Sem faturas para mostrar"))
		fmt.Println(err)
		return
	}
	bot.Send(tgbotapi.NewMessage(*chatID, fmt.Sprintf("Foram localizadas essas faturas para o mês %d:", mes)))
	for _, item := range faturas {
		var id = int64(item.ID)
		BotoesFatura(bot, chatID, id, models.ModelaFatura(&item), paga)
	}

}
