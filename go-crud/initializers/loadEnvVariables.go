package initializers

import "github.com/joho/godotenv"

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
