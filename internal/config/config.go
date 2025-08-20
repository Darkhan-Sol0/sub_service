package config

import (
	"log"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	once   sync.Once
	config Config
)

type (
	ServerConfig struct {
		ServerHTTP   `yaml:"server_http"`
		DatabasePG   `yaml:"database"`
		LoggerConfig `yaml:"logger"`
	}

	LoggerConfig struct {
		LogLevel string `yaml:"log_level"`
		LogOut   string `yaml:"log_out"`
	}

	ServerHTTP struct {
		Address     string        `yaml:"address"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
	}

	DatabasePG struct {
		Env      string `yaml:"database_env"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database_name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	Config interface {
		GetLogLevel() string
		GetLogOut() string

		GetAddress() string
		GetIdleTime() time.Duration

		GetDBEnv() string
		GetDBPort() string
		GetDBHost() string
		GetDBDatabase() string
		GetDBUsername() string
		GetDBPassword() string
	}
)

func LoadConfig() Config {
	once.Do(func() {
		config = &ServerConfig{}
		configPath := "config/config.yaml"
		if err := cleanenv.ReadConfig(configPath, config); err != nil {
			log.Fatalf("error read config %s: %v", configPath, err)
		}
	})
	return config
}

func (s *ServerConfig) GetAddress() string {
	return s.ServerHTTP.Address
}

func (s *ServerConfig) GetIdleTime() time.Duration {
	return s.ServerHTTP.IdleTimeout
}

func (s *ServerConfig) GetDBEnv() string {
	return s.DatabasePG.Env
}

func (s *ServerConfig) GetDBPort() string {
	return s.DatabasePG.Port
}

func (s *ServerConfig) GetDBHost() string {
	return s.DatabasePG.Host
}

func (s *ServerConfig) GetDBDatabase() string {
	return s.DatabasePG.Database
}

func (s *ServerConfig) GetDBUsername() string {
	return s.DatabasePG.Username
}

func (s *ServerConfig) GetDBPassword() string {
	return s.DatabasePG.Password
}

func (s *ServerConfig) GetLogLevel() string {
	return s.LoggerConfig.LogLevel
}

func (s *ServerConfig) GetLogOut() string {
	return s.LoggerConfig.LogOut
}
