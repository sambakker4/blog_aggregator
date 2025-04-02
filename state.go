package main

import (
	"github.com/sambakker4/blog_aggregator/internal/config"
	"errors"
	"fmt"
)

type state struct {
	config *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Login handler expects argument <username>")
	}
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("User:", cmd.args[0], "has been set as user")
	return nil
}
