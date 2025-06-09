package notificacao

import (
	"fmt"
	"telegram/funcionalidades/fatura"
	"telegram/models"
	"telegram/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)

func NotificaFaturaAvencerDoDia(bot *tgbotapi.BotAPI) {
	var chatid = int64(552380571)
	c := cron.New()
	// Agendando execução diária às 12:00
	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Executando rotina")
		faturas, err := repository.GetFaturasVencidasNoMesNaoPagas()
		if err != nil {
			fmt.Println(err)
		}
		bot.Send(tgbotapi.NewMessage(chatid, "‼️ ROTNA DE NOTIFICAÇÃO! ‼️"))
		bot.Send(tgbotapi.NewMessage(chatid, "Faturas com vencimento no dia de hoje ou vencidas!!"))
		for _, item := range *faturas {

			var id = int64(item.ID)
			fatura.BotoesFatura(bot, &chatid, id, models.ModelaFatura(&item), false)
		}
	})

	c.Start()

}
