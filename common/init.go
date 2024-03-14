package common

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	Port       string `default:":9170"`
	RootURL    string `split_words:"true" default:"/api"`
	NameApp    string `split_words:"true" default:"Backend Test"`
	LogLevel   string `split_words:"true" default:"debug"`
	ConfigFile string `split_words:"true" default:"config/config_prod.json"`
}

type fileConfig struct {
	Database struct {
		Postgre struct {
			Name     string `json:"name"`
			Host     string `json:"host"`
			Port     int    `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
			SSLMode  string `json:"sslmode"`
			LogMode  bool   `json:"logMode"`
			Schema   string `json:"schema"`
		} `json:"postgre"`
	} `json:"database"`
}

var Config Configuration
var FileConfig fileConfig

func InitConfig() {
	err := envconfig.Process("PC", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.WithField("config", Config).Info("Config successfully loaded")

	file, err := os.ReadFile(Config.ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal([]byte(file), &FileConfig)
}
