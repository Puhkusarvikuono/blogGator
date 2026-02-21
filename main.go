package main

import (
	"fmt"
	"github.com/Puhkusarvikuono/blogGator/internal/config"
	"log"
)

func main() {
	cfg, err := config.ReadConfigJson()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(cfg)

	err = cfg.SetUser("puhku")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}


	cfg, err = config.ReadConfigJson()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(cfg)

}
