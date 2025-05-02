package sessao

import (
	"fmt"
	"telegram/models"
)

func chaveSessao(userID int64, tipo models.Tipo) string {
	return fmt.Sprintf("%s:%s", fmt.Sprintf("%d", userID), tipo)
}

func CriarSessao(userID *int64, tipo models.Tipo) *models.SessaoFatura {
	chave := chaveSessao(*userID, tipo)
	if models.Sessoes[chave] == nil {
		models.Sessoes[chave] = &models.SessaoFatura{
			EtapaAtual: models.EtapaValor,
			Tipo:       tipo,
		}
	}
	return models.Sessoes[chave]
}

func VerificaSessao(userID *int64, tipo models.Tipo) (*models.SessaoFatura, bool) {
	chave := chaveSessao(*userID, tipo)
	sessao, existe := models.Sessoes[chave]
	if !existe {
		return nil, false
	}
	if sessao.Tipo != tipo {
		return nil, false
	}
	//fmt.Println(sessao.Fatura.Descricao)
	return sessao, true
}

func RemoveSessao(userID *int64, tipo models.Tipo) {
	sessao, existe := VerificaSessao(userID, tipo)
	if existe && sessao.Tipo == tipo {
		chave := chaveSessao(*userID, tipo)
		delete(models.Sessoes, chave)
		return
	}
	fmt.Println(fmt.Sprint("sessão não removida: %w", sessao))
}

func VerificaContinuidadeInserirFatura(userID *int64) bool {
	sec, b := VerificaSessao(userID, models.TipoInserir)
	if b && sec != nil {
		return true
	}
	return false
}
