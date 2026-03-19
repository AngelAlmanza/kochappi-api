package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL         string
	JWTSecret           string
	JWTAccessExpiryMin  int
	JWTRefreshExpiryDay int
	Port                string
	Env                 string
	LogLevel            string
	OTPExpiryMinutes    int
}

func Load() *Config {
	_ = godotenv.Load("config/.env.local")

	return &Config{
		DatabaseURL:         getEnv("DATABASE_URL", "postgresql://kochappi:password@localhost:5432/kochappi_dev"),
		JWTSecret:           getEnv("JWT_SECRET", "dev-secret-change-in-production"),
		JWTAccessExpiryMin:  getEnvAsInt("JWT_ACCESS_EXPIRY_MINUTES", 15),
		JWTRefreshExpiryDay: getEnvAsInt("JWT_REFRESH_EXPIRY_DAYS", 7),
		Port:                getEnv("PORT", "8081"),
		Env:                 getEnv("ENV", "development"),
		LogLevel:            getEnv("LOG_LEVEL", "debug"),
		OTPExpiryMinutes:    getEnvAsInt("OTP_EXPIRY_MINUTES", 10),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return fallback
}
