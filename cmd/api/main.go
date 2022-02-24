package main

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"storage/configs"
)

const(
	configPath = "./configs/config.yml"
)

var conf configs.Conf

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{})

	log.Info("initialize config")
	bb, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bb, &conf)
	if err != nil {
		log.Fatal(err)
	}
}


func main()  {
	err := app()
	if err != nil {
		log.Fatalln(err)
	}
}

