package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		fmt.Println("Reset takes no arguments")
	}
	err := s.db.Reset(context.Background())
	if err != nil {
		fmt.Println("Reset unsuccessful")
		os.Exit(1)
	}
	err = s.cfg.SetUser("")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return nil
}
