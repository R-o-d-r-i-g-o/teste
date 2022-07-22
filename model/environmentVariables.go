package model

import (
	controFunc "gitTest/controller"

	v6 "github.com/caarlos0/env/v6"
)

type EnvironmentVariables struct {
	ClientId    string `env:"CLIENT_ID"`
	ClientToken string `env:"CLIENT_TOKEN"`
}

func (env *EnvironmentVariables) Load() {
	env.setWithEnvFile()
	env.cleanEnvironmentVariables()
}

func (env *EnvironmentVariables) setWithEnvFile() {
	v6.Parse(env)
}

func (env *EnvironmentVariables) cleanEnvironmentVariables() {
	env.ClientId = controFunc.LetOnlyNumbers(env.ClientId)
	env.ClientToken = controFunc.RevomeSpecialChars(env.ClientToken)
}
