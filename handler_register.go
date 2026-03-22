package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Puhkusarvikuono/blogGator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("the handler expects a single argument, the username")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		fmt.Println("User already exists")
		os.Exit(1)
	}

	if s.cfg.CurrentUserName == user.Name {
		fmt.Println("Current user already logged in")
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		fmt.Println("Error creating user")
		os.Exit(1)
	}

	fmt.Println("User created successfully!")
	fmt.Println(user)

	return nil
}
