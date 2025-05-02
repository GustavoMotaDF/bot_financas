package handlers

import (
	"strings"
	"telegram/funcionalidades/fatura"
	"telegram/handlers/menu"
	"telegram/models"
	"telegram/sessao"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Zero int64 = 0

func Start(bot *tgbotapi.BotAPI, update *tgbotapi.Update, userID *int64, chatID *int64, msg string, msgID *int) {

	if msg == "/inserirFatura" {
		sessao.CriarSessao(userID, models.TipoInserir)
		bot.Send(tgbotapi.NewMessage(*chatID, "**************"))
		bot.Send(tgbotapi.NewMessage(*chatID, "Digite o valor da fatura (ex: 123.45):"))
		return
	}
	if strings.HasPrefix(msg, "/consultar_") {
		Consultar(bot, update, userID, chatID, msg, msgID)
		return
	}
	if strings.HasPrefix(msg, "/delete_") {
		fatura.DeleteFatura(bot, userID, chatID, msg, msgID)
		return
	}
	if strings.HasPrefix(msg, "/pagar_") {
		fatura.PagarFatura(bot, userID, chatID, msg, msgID)
		return
	}
	if strings.HasPrefix(msg, "/relatorios_") {
		Relatorios(bot, update, userID, chatID, msg, msgID)
		return
	}

	if sessao.VerificaContinuidadeInserirFatura(userID) {
		fatura.Inserirfatura(bot, userID, chatID, msg)
		return
	}
	menu.ShowMenu(bot, userID)

}
