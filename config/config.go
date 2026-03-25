package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	AppConfig *Config
)

type Config struct {
	AppPort string
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
	JWTSecret string
	JWTExpired string
	RefreshToken string
	AllowOrigins string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	AppConfig = &Config{
		AppPort: getEnv("PORT", "3000"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "3306"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName: getEnv("DB_NAME", "himaose"),
		JWTSecret: getEnv("JWT_SECRET", "super_secret"),
		JWTExpired: getEnv("JWT_EXPIRED", "6h"),
		RefreshToken: getEnv("REFRESH_TOKEN_EXPIRED", "24h"),
		AllowOrigins: getEnv("ALLOW_ORIGINS", ""),
	}
}

func getEnv(key string, fallback string) string {
	value, exist := os.LookupEnv(key)
	if exist {
		return value
	} else {
		return fallback
	}
}

func ConnectDB() {
	cfg := AppConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
        cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}