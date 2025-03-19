package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

//type AuthService struct {
//	Host string
//	Port string
//}

type DBConf struct {
	Host string `env:"DB_HOST"`
	Port string `env:"DB_PORT"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
	Name string `env:"DB_NAME"`
}

//type Redis struct {
//	Host string
//	Port string
//}

type ServerConf struct {
	Host string `env:"SERVER_HOST"`
	Port string `env:"SERVER_PORT"`
}

type Config struct {
	DB     DBConf
	Server ServerConf
}

var globalConfig *Config

func GetConfig() *Config {
	if globalConfig == nil {
		globalConfig = &Config{}
	}

	return globalConfig
}

func ParseConfig(path ...string) {
	if len(path) != 0 {
		//если передан путь к файлу, то парсим файл, если нет, то читаем переменные окружения
		err := cleanenv.ReadConfig(path[0], GetConfig())
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	err := cleanenv.ReadEnv(GetConfig())
	if err != nil {
		log.Fatal(err)
	}

	return
}
