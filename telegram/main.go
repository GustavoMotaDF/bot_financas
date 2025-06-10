package main

import (
	"log"
	"telegram/config"
	"telegram/config/auth"
	"telegram/config/bd"
	"telegram/funcionalidades/notificacao"
	"telegram/handlers"
	"telegram/models"
	"telegram/sessao"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bd.ConnectDB()
	config.LoadConfig()

	bot, err := tgbotapi.NewBotAPI(config.AppConfig.ApiKey)
	if err != nil {
		log.Panic(err)
	}

	// Deletando o webhook com a API
	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})

	if err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			userID := handlers.ExtrairUserID(&update)
			chatID := handlers.ExtrairChatID(&update)
			sessao.CriarSessao(userID, models.TipoBasica)
			if auth.VerificaPermissao(userID) {
				msg := handlers.ExtrairMensagem(&update)
				msgID := handlers.ExtrairMensagemID(&update)
				handlers.Start(bot, &update, userID, chatID, msg, msgID)
			} else {
				bot.Send(tgbotapi.NewMessage(*chatID, " ⛔ Este é um bot privado! ⛔"))
			}
			sessao.RemoveSessao(userID, models.TipoBasica)
		}
	}()
	go func() {
		notificacao.NotificaFaturaAvencerDoDia(bot)
	}()
	select {}

}
