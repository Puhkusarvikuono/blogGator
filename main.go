package main

import (
	"fmt"
	"github.com/Puhkusarvikuono/blogGator/internal/config"
	"errors"
	"os"
	"log"
)

type state struct {
	ConfigPtr			*config.Config
}

type command struct {
	Name		string
	Args		[]string
}

type commands struct {
	commands	map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
}


func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("The handler expects a single argument, the username.")
	}

	login := cmd.Args[0]
	
	err := s.ConfigPtr.SetUser(login)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("The user has been set:", login)
	return nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: Not enough arguments.")
		os.Exit(1)
	}

	newCommand := command{
		Name: args[1],
		Args: []string{},
	}

	if len(args) > 2 {
		newCommand.Args = args[2:]
	}
	
	cfg, err := config.ReadConfigJson()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	newState := &state{
		ConfigPtr:	&cfg,
	}

	
	newCommands :=	map[string]func(*state, command) error {
		"login": handlerLogin,
	}

	if commandFunc, exists := newCommands[newCommand.Name]; exists {
		err := commandFunc(newState, newCommand)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	} else {
			fmt.Println("Invalid command")
			os.Exit(1)
	}
}
