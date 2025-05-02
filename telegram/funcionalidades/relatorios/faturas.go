package relatorios

import (
	"fmt"
	"telegram/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RetornaValorAPagarMes(mes int, ano int) {

}
func RetornaValorPagoMes(mes int) {

}
func RelatorioMes(mes int, ano int, bot *tgbotapi.BotAPI, userID *int64, chatID *int64) {

	faturasPagas, err := repository.GetFaturasMes(mes, ano, true)
	faturasNaoPagas, errr := repository.GetFaturasMes(mes, ano, false)

	if len(faturasPagas) == 0 && len(faturasNaoPagas) == 0 || err != nil && errr != nil {
		bot.Send(tgbotapi.NewMessage(*chatID, "‼️ Sem faturas para somar!"))
		fmt.Println(err)
		return
	}

	totalPagas := SomaFaturas(faturasPagas)
	totalNaoPagas := SomaFaturas(faturasNaoPagas)

	strPagas := FaturasEmLinhas(faturasPagas)
	strNPagas := FaturasEmLinhas(faturasNaoPagas)

	bot.Send(tgbotapi.NewMessage(*chatID, fmt.Sprintf("Faturas Pagas : %s\nValor total: R$%0.2f", strPagas, totalPagas)))
	bot.Send(tgbotapi.NewMessage(*chatID, fmt.Sprintf("Faturas Não Pagas : %s\nValor total: R$%0.2f", strNPagas, totalNaoPagas)))

	total := totalPagas + totalNaoPagas

	bot.Send(tgbotapi.NewMessage(*chatID, fmt.Sprintf("Valor total para o mês %d\nR$%0.2f", mes, total)))

}
