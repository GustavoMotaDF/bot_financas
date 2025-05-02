package relatorios

import (
	"fmt"
	"telegram/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ListarRelatorios(bot *tgbotapi.BotAPI, chatID *int64) {
	msg := tgbotapi.NewMessage(*chatID, "Escolha uma opção:")
	buttons := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("Por Mês", "/relatorios_total_mes"),
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons)
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}
func RetornaListaMesesAnos(bot *tgbotapi.BotAPI, chatID *int64) {
	msg := tgbotapi.NewMessage(*chatID, "Selecione o mês/ano:")

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 1; i <= 12; i++ {
		mesBtn := tgbotapi.NewInlineKeyboardButtonData(

			fmt.Sprintf("%02d/2025", i),
			///consultar_faturas_mes_4_2025_paga_true // false
			fmt.Sprintf("/relatorios_total_mes_%d_2025", i),
		)

		// Agrupar por 3 botões por linha
		if (i-1)%4 == 0 {
			rows = append(rows, []tgbotapi.InlineKeyboardButton{})
		}
		rows[len(rows)-1] = append(rows[len(rows)-1], mesBtn)
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)

	bot.Send(msg)

}

func SomaFaturas(faturasPagas []models.Fatura) float64 {
	var valorFaturas []float64
	var total float64

	for _, fatura := range faturasPagas {
		valorFaturas = append(valorFaturas, fatura.Valor)
	}

	for _, valor := range valorFaturas {
		total = total + valor
	}

	return total
}

func FaturasEmLinhas(faturas []models.Fatura) string {
	var strFaturas []string
	for _, item := range faturas {
		strFaturas = append(strFaturas, models.ModelaFaturaEmUmaLinha(&item))
	}

	var msg string

	for _, item := range strFaturas {
		msg = msg + "\n" + item
		msg = msg + "\n__________________________"
	}

	return msg
}
