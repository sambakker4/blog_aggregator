package main

import (
	"errors"
	"fmt"
	"context"
	"github.com/google/uuid"
	"github.com/sambakker4/blog_aggregator/internal/database"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Login handler expects argument <username>")
	}

	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err != nil {
		return errors.New("User does not exist")
	} 
	
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("User:", cmd.args[0], "has been set as user")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Register expects argument <name>")
	}

	name := cmd.args[0]
	if _, err := s.db.GetUser(context.Background(), name); err == nil {
		return errors.New("User already exists")
	}

	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		Name: name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	fmt.Println("User", name, "was created")
	if err != nil {
		return err
	}

	err = s.config.SetUser(name)
	if err != nil {
		return err
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Database reset successful")
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())	
	if err != nil {
		return err
	}

	currentUser := s.config.CurrentUserName

	for _, user := range users {
		fmt.Print(" * ", user)
		if user == currentUser {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

	return nil
}
