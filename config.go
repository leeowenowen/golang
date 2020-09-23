package main

import (
	"log"

	"github.com/Unknwon/goconfig"
)

func loadConfig(env string) {
	cfg, err := goconfig.LoadConfigFile("dev-config.ini")
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}
	value, error := cfg.GetValue("redis", "host")
	if error != nil {
		log.Printf("failed to get key host", error)
		return
	}
	log.Printf("get value by key redis.host %s", value)

}

func main() {
	//load config
	loadConfig("dev")
}
