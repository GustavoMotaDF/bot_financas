package handlers

import (
	"strings"
	"telegram/funcionalidades/relatorios"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Relatorios(bot *tgbotapi.BotAPI, update *tgbotapi.Update, userID *int64, chatID *int64, msg string, msgID *int) {
	if msg == "/relatorios_listar" {
		relatorios.ListarRelatorios(bot, chatID)
		return
	}
	if msg == "/relatorios_total_mes" {
		relatorios.RetornaListaMesesAnos(bot, chatID)
		return
	}
	///relatorios_valor_total_mes_4_2025
	if strings.HasPrefix(msg, "/relatorios_total_mes_") {

		mesAno := strings.Split(msg, "_")
		var ano = ConvertParaInt(mesAno[4])
		var mes = ConvertParaInt(mesAno[3])
		relatorios.RelatorioMes(mes, ano, bot, userID, chatID)
		return
	}

}
