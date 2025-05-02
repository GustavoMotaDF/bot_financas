package menu

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowMenu(bot *tgbotapi.BotAPI, chatID *int64) {
	var rows [][]tgbotapi.InlineKeyboardButton
	msg := tgbotapi.NewMessage(*chatID, "Para começar, selecione uma opção: ")

	btnInserir := tgbotapi.NewInlineKeyboardButtonData("Inserir uma Fatura", "/inserirFatura")
	btnConsultar := tgbotapi.NewInlineKeyboardButtonData("Consultar Faturas", "/consultar_faturas")
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Relatórios", "/relatorios_listar")

	rows = append(rows, []tgbotapi.InlineKeyboardButton{})

	rows[len(rows)-1] = append(rows[len(rows)-1], btnInserir)
	rows[len(rows)-1] = append(rows[len(rows)-1], btnConsultar)
	rows[len(rows)-1] = append(rows[len(rows)-1], btn1)

	// rows = append(rows, []tgbotapi.InlineKeyboardButton{})

	// btn2 := tgbotapi.NewInlineKeyboardButtonData("Consultar Faturas", "/consultar_faturas")
	// btn3 := tgbotapi.NewInlineKeyboardButtonData("Consultar Faturas", "/consultar_faturas")

	// rows[len(rows)-1] = append(rows[len(rows)-1], btn2)
	// rows[len(rows)-1] = append(rows[len(rows)-1], btn3)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)

	bot.Send(msg)

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
