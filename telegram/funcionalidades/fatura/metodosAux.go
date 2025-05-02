package fatura

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ConvertParaInt(valor string) int {
	mes, err := strconv.Atoi(valor)
	if err != nil {
		fmt.Println("Erro ao converter em string")
		return -1
	}
	return mes
}

func RetornaListaMesesAnos(bot *tgbotapi.BotAPI, chatID *int64, paga bool) {
	msg := tgbotapi.NewMessage(*chatID, "Selecione o mês/ano:")

	var rows [][]tgbotapi.InlineKeyboardButton
	defaultbtn := tgbotapi.NewInlineKeyboardButtonData("Ver todas faturas", fmt.Sprintf("/consultar_getAllfaturas_pagas_%t", paga))
	// Adiciona botões em linhas de 3 em 3, por exemplo
	for i := 1; i <= 12; i++ {
		mesBtn := tgbotapi.NewInlineKeyboardButtonData(

			fmt.Sprintf("%02d/2025", i),
			///consultar_faturas_mes_4_2025_paga_true // false
			fmt.Sprintf("/consultar_faturas_mes_%02d_2025_%t", i, paga),
		)

		// Agrupar por 3 botões por linha
		if (i-1)%4 == 0 {
			rows = append(rows, []tgbotapi.InlineKeyboardButton{})
		}
		rows[len(rows)-1] = append(rows[len(rows)-1], mesBtn)
	}
	rows = append(rows, tgbotapi.NewInlineKeyboardRow(defaultbtn))
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)

	bot.Send(msg)

}

func EscolhaFaturaPagaNaoPaga(bot *tgbotapi.BotAPI, chatID *int64) {
	msg := tgbotapi.NewMessage(*chatID, "Escolha uma opção:")
	buttons := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("Faturas pagas", "/consultar_faturas_pagas_true"),
		tgbotapi.NewInlineKeyboardButtonData("Faturas Não pagas", "/consultar_faturas_pagas_false"),
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons) // isso já cria uma única linha
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func BotoesFatura(bot *tgbotapi.BotAPI, chatID *int64, faturaID int64, fatura string, paga bool) {
	deleteButton := tgbotapi.NewInlineKeyboardButtonData("🗑️ Deletar", fmt.Sprintf("/delete_%d", faturaID))
	pagarbtn := tgbotapi.NewInlineKeyboardButtonData("💲 Pagar", fmt.Sprintf("/pagar_%d", faturaID))
	msg := tgbotapi.NewMessage(*chatID, fatura)
	if !paga {
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(deleteButton),
			tgbotapi.NewInlineKeyboardRow(pagarbtn),
		)
		msg.ReplyMarkup = keyboard
	}
	bot.Send(msg)

}
