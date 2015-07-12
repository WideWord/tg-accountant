package config

import(
	"code.google.com/p/gcfg"
	"log"
)

type Config struct {
	Bot struct {
		Token string
	}
	Redis struct {
		Host string
		Password string
	}
}

var config Config

var readed bool = false

func Read() {
	err := gcfg.ReadFileInto(&config, "accountant.gcfg")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Get() Config {
	return config
}
