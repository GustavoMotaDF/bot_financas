package models

type SessaoFatura struct {
	EtapaAtual Etapa
	Fatura     Fatura
	Tipo       Tipo
}

var Sessoes = make(map[string]*SessaoFatura)

type Tipo string

const (
	TipoBasica    Tipo = "0"
	TipoInserir   Tipo = "1"
	TipoConsultar Tipo = "2"
	TipoDeletar   Tipo = "3"
	TipoAlterar   Tipo = "4"
)
