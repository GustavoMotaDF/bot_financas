package handlers

import (
	"strings"
	"telegram/funcionalidades/fatura"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Consultar(bot *tgbotapi.BotAPI, update *tgbotapi.Update, userID *int64, chatID *int64, msg string, msgID *int) {

	if msg == "/consultar_faturas" {
		// retorna botão perguntando pagas ou não pagas
		fatura.EscolhaFaturaPagaNaoPaga(bot, chatID)
		return
	}
	///consultar_faturas_pagas_true // false
	if strings.HasPrefix(msg, "/consultar_faturas_pagas_") {
		mss := strings.Split(msg, "_")
		valor := mss[3]
		paga, err := ConvertParaBool(valor)
		if err != nil {
			return
		}
		fatura.RetornaListaMesesAnos(bot, chatID, paga)
		return
	}
	///consultar_faturas_mes_4_2025_paga_true // false
	if strings.HasPrefix(msg, "/consultar_faturas_mes_") {
		mss := strings.Split(msg, "_")
		valor := mss[5]
		paga, err := ConvertParaBool(valor)
		if err != nil {
			return
		}
		fatura.GetFaturasMes(bot, chatID, msg, paga)
		return
	}
	///consultar_getAllfaturas_pagas_true // false
	if strings.HasPrefix(msg, "/consultar_getAllfaturas_pagas_") {
		mss := strings.Split(msg, "_")
		valor := mss[3]
		paga, err := ConvertParaBool(valor)
		if err != nil {
			return
		}
		fatura.GetAllFaturas(bot, chatID, paga)
		return
	}

}
