package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/zn/http-rest-api/internal/app/apiserver"
	"log"
	"os"
)

var(
	configPath string
)

func init(){
	flag.StringVar(&configPath, "config-path", fmt.Sprintf("configs%capiserver.toml", os.PathSeparator), "path to config file")
}

func main(){
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil{
		log.Fatal(err)
	}

	s:= apiserver.New(config)
	if err := s.Start(); err != nil{
		log.Fatal(err)
	}
}