package auth

import (
	"telegram/config"
)

func VerificaPermissao(userID *int64) bool {

	for _, item := range config.AppConfig.UserID {
		if item == *userID {
			return true
		}
	}
	return false
}
