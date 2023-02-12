package environment

import (
	"os"
	"sync"
)

type Config struct {
	Env map[string]string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	if instance == nil {
		once.Do(func() {
			c := &Config{}
			c.getEnvVars()
			instance = c
		})
	}
	return instance
}

func (c *Config) getEnvVars() {
	if token := os.Getenv("KREW_GITHUB_TOKEN"); token != "" {
		c.Env["KREW_GITHUB_TOKEN"] = token
	}
}
