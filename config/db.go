package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erreur de chargement du fichier .env")
	}

	// Construire la chaîne de connexion
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
        DSN:                  dsn,
        PreferSimpleProtocol: true, // ✅ Désactive les prepared statements automatiques
    }), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base :", err)
	}

	fmt.Println("✅ Connexion à PostgreSQL réussie !")
	DB = db

}
