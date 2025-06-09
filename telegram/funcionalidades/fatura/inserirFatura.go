package fatura

import (
	"fmt"
	"strconv"
	"strings"
	"telegram/handlers/menu"
	models "telegram/models"
	"telegram/repository"
	sessao "telegram/sessao"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Inserirfatura(bot *tgbotapi.BotAPI, userID *int64, chatID *int64, text string) {
	sessaoAtual, boole := sessao.VerificaSessao(userID, models.TipoInserir)

	if !boole {
		return
	}

	switch sessaoAtual.EtapaAtual {
	case models.EtapaValor:

		valorStr := text
		valorStr = strings.Replace(valorStr, ",", ".", 1)
		valorFloat, _ := strconv.ParseFloat(valorStr, 64)
		if valorFloat < 0 {
			bot.Send(tgbotapi.NewMessage(*chatID, "Valor não pode ser negativo!"))
			bot.Send(tgbotapi.NewMessage(*chatID, "Digite o valor da fatura (ex: 123.45):"))
			return
		}
		sessaoAtual.Fatura.Valor = valorFloat
		sessaoAtual.EtapaAtual = models.EtapaDescricao
		bot.Send(tgbotapi.NewMessage(*chatID, "Digite a descrição da fatura:"))

	case models.EtapaDescricao:
		sessaoAtual.Fatura.Descricao = text
		sessaoAtual.EtapaAtual = models.EtapaVencimento
		bot.Send(tgbotapi.NewMessage(*chatID, "Digite a data de vencimento (DD/MM/AAAA):"))
	case models.EtapaVencimento:
		vencimento, err := time.Parse("02/01/2006", text)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(*chatID, "Data inválida. Use o formato DD/MM/AAAA."))
			return
		}
		if vencimento.Before(time.Now()) {
			bot.Send(tgbotapi.NewMessage(*chatID, "A data informada deve ser igual ou maior a atual!"))
			return
		}
		sessaoAtual.Fatura.Vencimento = vencimento
		sessaoAtual.Fatura.UserID = *userID

		mensagem := repository.SaveFatura(&sessaoAtual.Fatura)

		bot.Send(tgbotapi.NewMessage(*chatID, fmt.Sprintln(mensagem)))
		bot.Send(tgbotapi.NewMessage(*chatID, "**************"))
		menu.ShowMenu(bot, chatID)
		// Finaliza a sessão
		sessao.RemoveSessao(userID, models.TipoInserir)
		return
	}

}
