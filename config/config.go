package config

import (
	"github.com/IlyaKharitonov/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type DBConf struct {
	Host string `env:"DB_HOST"`
	Port string `env:"DB_PORT"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
	Name string `env:"DB_NAME"`
}

type ServerConf struct {
	Host string `env:"SERVER_HOST"`
	Port string `env:"SERVER_PORT"`
}

type configType struct {
	DB     DBConf
	Server ServerConf
	Logger logger.LoggerConf
}

//
//func GetConfig() *Config {
//	if globalConfig == nil {
//		globalConfig = &Config{}
//	}
//
//	return globalConfig
//}

func ParseConfig(path ...string) *configType {
	var config = &configType{}

	if len(path) != 0 {
		//если передан путь к файлу, то парсим файл, если нет, то читаем переменные окружения
		err := cleanenv.ReadConfig(path[0], config)
		if err != nil {
			log.Fatal(err)
		}

		return config
	}

	err := cleanenv.ReadEnv(config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
