package main

import (
	"gitTest/model"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	var envVars model.EnvironmentVariables = model.EnvironmentVariables{}

	envVars.Load()
}
