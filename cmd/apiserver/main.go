package main

import (
	"flag"
	"go_api/internal/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init(){
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	//создаем конфиг
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	//создаем сервер с данными конфига
	s := apiserver.New(config)

	// fmt.Println(s)

	//стартуем сервер
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
