package main

import (
	"context"
	"os"
	"fmt"
)


func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("No users to list.")
		os.Exit(1)
	}
	for _, user := range users {
		nameOutput := user.Name
		if s.cfg.CurrentUserName == user.Name {
			nameOutput = fmt.Sprintf("%s (current)", user.Name)
		}
		fmt.Printf("* %s\n", nameOutput) 
	}
	return nil
}
