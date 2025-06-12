package repository

import (
	"fmt"
	"telegram/config"
	"telegram/config/bd"
	"telegram/models"
	"time"
)

// Função para salvar um evento no banco
func SaveFatura(fatura *models.Fatura) string {
	err := bd.DB.Create(fatura).Error
	if err != nil {
		return fmt.Errorf("‼️ Erro ao salvar fatura: %w", err).Error()
	} else {
		return fmt.Sprint("✅ Fatura ", fatura.Descricao, " salva com sucesso! ")
	}

}

func GetAllFaturas(paga bool) ([]models.Fatura, error) {
	var faturas []models.Fatura

	err := bd.DB.Model(&faturas).
		Where("user_id IN ?", config.AppConfig.UserID).
		Where("deleted_at IS NULL").
		Where("paga = ?", paga).
		Find(&faturas).Error
	return faturas, err
}

func GetFaturasMes(mes int, ano int, paga bool) ([]models.Fatura, error) {
	var faturas []models.Fatura
	err := bd.DB.Model(&faturas).
		Where("user_id IN ?", config.AppConfig.UserID).
		Where("strftime('%m', vencimento) = ? AND strftime('%Y', vencimento) = ?", fmt.Sprintf("%02d", mes), fmt.Sprintf("%d", ano)).
		Where("deleted_at IS NULL").
		Where("paga = ?", paga).
		Find(&faturas).Error
	return faturas, err
}

func PagarFatura(id *int) error {
	var faturas []models.Fatura

	return bd.DB.Model(&faturas).
		Where("user_id IN ?", config.AppConfig.UserID).
		Where("id = ?", id).Update("paga", true).Error
}

func DeleteFatura(fatura *int64) error {
	var faturas []models.Fatura

	return bd.DB.Model(&faturas).
		Where("user_id IN ?", config.AppConfig.UserID).
		Where("id = ?", *fatura).Update("deleted_at", time.Now()).Error
}

func GetFaturasVencidasNoMesNaoPagas() ([]models.Fatura, error) {
	var faturas []models.Fatura
	err := bd.DB.Model(&faturas).
		Where("strftime('%Y-%m-%d', vencimento) >= strftime('%Y-%m', 'now') || '-01'").
		Where("strftime('%Y-%m-%d', vencimento) <= strftime('%Y-%m-%d', 'now')").
		Where("paga = ?", false).
		Find(&faturas).Error

	return faturas, err
}
