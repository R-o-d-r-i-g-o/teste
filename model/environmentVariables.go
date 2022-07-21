package model

import (
	v6 "github.com/caarlos0/env/v6"
)

type EnvironmentVariables struct {
	ClientId    string `env:"CLIENT_ID"`
	ClientToken string `env:"CLIENT_TOKEN"`
}

func (env *EnvironmentVariables) Load() {
	env.setWithEnvFile()
}

func (env *EnvironmentVariables) setWithEnvFile() {
	v6.Parse(env)
}
