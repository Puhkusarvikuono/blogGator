package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("the handler expects a single argument, the username")
	}

	login := cmd.Args[0]

	if s.cfg.CurrentUserName == login {
		fmt.Println("Current user already logged in")
		os.Exit(1)
	}

	_, err := s.db.GetUser(context.Background(), login)
	if err != nil {
		fmt.Println("The user is not registered")
		os.Exit(1)
	}

	err = s.cfg.SetUser(login)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("The user has been set:", login)
	return nil
}
