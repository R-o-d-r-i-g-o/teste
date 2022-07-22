package main

import (
	"fmt"
	"gitTest/model"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	var envVars model.EnvironmentVariables = model.EnvironmentVariables{}

	envVars.Load()

	fmt.Println(envVars)
}
