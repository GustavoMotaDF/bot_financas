package notificacao

import (
	"fmt"
	"telegram/config"
	"telegram/funcionalidades/fatura"
	"telegram/models"
	"telegram/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)

func NotificaFaturaAvencerDoDia(bot *tgbotapi.BotAPI) {
	c := cron.New()
	// Agendando execução diária às 12:00
	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Executando rotina")
		faturas, err := repository.GetFaturasVencidasNoMesNaoPagas()
		if err != nil {
			fmt.Println(err)
		}
		if faturas == nil {
			fmt.Println("Notificando usuários")
			//notificando todos os usuários informados na variavel de ambiente TELEGRAM_USER_IDS
			for _, id := range config.AppConfig.UserID {

				bot.Send(tgbotapi.NewMessage(id, "‼️ ROTNA DE NOTIFICAÇÃO! ‼️"))
				bot.Send(tgbotapi.NewMessage(id, "Faturas com vencimento no dia de hoje ou vencidas!!"))
				for _, item := range *faturas {

					var id = int64(item.ID)
					fatura.BotoesFatura(bot, &id, id, models.ModelaFatura(&item), false)
				}
			}
		} else {
			fmt.Println("Sem faturas para notificar")
		}
	})

	c.Start()

}
