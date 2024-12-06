package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

}
