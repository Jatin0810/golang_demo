package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env        string `yaml:"env"`
	Db_Url string `yaml: "db_url"`
	HTTPServer `yaml:"http_server"`
}

func CongifInit() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")


	if configPath == ""{

		flags := flag.String("config","", "path to the config path")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}


	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Printf("config file does not exist %s", configPath)
	}

	var cfg Config

	readError := cleanenv.ReadConfig(configPath, &cfg)

	if readError != nil{
		log.Printf("can not read config file : %s", readError.Error())
	}


	return &cfg

}