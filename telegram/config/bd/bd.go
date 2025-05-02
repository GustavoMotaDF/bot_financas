package bd

import (
	"log"
	"telegram/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Função para conectar ao banco de dados SQLite
func ConnectDB() (*gorm.DB, error) {
	// Abrindo a conexão com o banco de dados SQLite
	var err error
	DB, err = gorm.Open(sqlite.Open("banco.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco: " + err.Error())
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso!")

	// Realizando o AutoMigrate para criar as tabelas automaticamente
	err = DB.AutoMigrate(&models.Fatura{})
	if err != nil {
		log.Fatalf("Erro ao realizar AutoMigrate: %v", err)
		return nil, err
	}
	log.Println("Migração do banco de dados concluída com sucesso!")

	return DB, nil
}
