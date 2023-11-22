package app

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

type config struct {
	Env             string
	ApplicationName string
	Db              struct {
		Host     string
		Port     string
		User     string
		Password string
		DbName   string `yaml:"dbName"`
		SslMode  bool
	}
}

var Config config

// loadConfig load all config of the application, config.yaml should be in same path of the binary
func loadConfig() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		path, _ := os.Getwd()
		log.WithField("path", path).Fatal("error while try to load config file", path)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatal("error while try to unmarshal the config file")
	}
}
