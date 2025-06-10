# 🤖 Bot de Finanças no Telegram

Este projeto é um bot de controle de contas a pagar pessoal desenvolvido em Go, que permite registrar e consultar faturas via Telegram.

## 🚀 Funcionalidades

- Inserir faturas
- Consultar por mês/ano
- Listar todas as faturas
- Marcar como paga
- Excluir
- Relatorio básico mensal, total pago e não pago

---

## 🧾 Pré-requisitos

- Go 1.23+ (se for rodar localmente)
- Docker (para uso com contêiner)
- Uma conta de bot no Telegram e seu token

---

## ⚙️ Uso

### 📄 Docker

Para uso via docker, pode-se usar a imagem (arm64 e amd64) gustavomota/bot_financas:0.0.1.0

#### Arquivo .env
Crie um arquivo .env para guardar as variaveis de ambiente necessárias:
- TELEGRAM_USER_IDS=xxxxxx,xxxxxxx,xxxxxx
- API_KEY=xxxxxxxxxxxxxxxx

Apos isso, rode:
``` 
docker run --env-file .env -v /path/to/bd:/app/config gustavomota/bot_financas:0.0.1.0' 
```
O volume é obrigatório para iniciar o container e persistir os dados via sqlite.

### 📄 Local
Baixe o projeto:
```
git clone https://github.com/GustavoMotaDF/bot_financas.git
```

Para baixar as dependencias execute: 
``` 
go run tidy 
```
ou
```
go mod tidy
```
Exporte as variavies de ambiente necessárias:
``` 
export TELEGRAM_USER_IDS="lista de ids" 
export API_KEY="sua api key" 
```
E para executar:

```
go run main.go
```

## Change Log
### 0.0.0.1 - alpha
Projeto base com as principais funcionalidades do mvp.
### 0.0.1.0 - Ajuste na validação da inserção de fatura
- 1) Validando valor negativo, impedindo inserir fatura com valor negativo
- 2) Validando datas, impedindo inserir fatura retroativa
### 0.0.2.0 - Ajuste de validação de datas, ao inserir fatura
- 1) BUG - Ajuste na validação de datas, até então não era possivel inserir fatura com data atual.