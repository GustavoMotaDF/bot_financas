package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Fatura struct {
	gorm.Model
	ID         int `gorm:"primaryKey"`
	Descricao  string
	Valor      float64
	Vencimento time.Time
	Paga       bool
	UserID     int64
}

func ModelaFatura(fatura *Fatura) string {
	var msg = fmt.Sprintf("ID: %d, \nDescrição: %s, \nValor: R$%.2f, \nVencimento: %s", fatura.ID, fatura.Descricao, fatura.Valor, fatura.Vencimento.Format("02/01/2006"))
	return msg
}

func ModelaFaturaEmUmaLinha(fatura *Fatura) string {
	var msg = fmt.Sprintf("Descrição: %s, Valor: R$%.2f, Vencimento: %s", fatura.Descricao, fatura.Valor, fatura.Vencimento.Format("02/01/2006"))
	return msg
}
