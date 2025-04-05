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

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")	
	if err != nil {
		return err
	}
	fmt.Println(*feed)
	return nil
}

func handlerNewFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("add feed requires a <name> and <url> arguments")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	user_id := user.ID
	
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user_id,
	})
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(
		context.Background(), database.CreateFeedFollowParams{	
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID:  feed.ID,
		})
	fmt.Println("User:", user.Name, "created feed called", feed.Name, "and is now following it")
	return nil 
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())	
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println("Name:",feed.Name)
		fmt.Println("URL:", feed.Url)
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println("Creator:", user.Name)
		fmt.Println("-----------------------------")
	}
	return nil
}


func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("follow commands expects argument <url>")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	
	_, err = s.db.CreateFeedFollow(
		context.Background(), database.CreateFeedFollowParams{	
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID:  feed.ID,
		})
	if err != nil {
		return err
	}	
	fmt.Println("User:", user.Name, "is now following:", feed.Name)
	fmt.Println()
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	
	fmt.Println("User:", user.Name, "is following:")
	for _, feed := range feeds {
		fmt.Print(" - '", feed.FeedName, "'\n")	
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("Unfollow command requires a <url> argument")
	}
	url := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,	
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("User:", user.Name, "successfully unfollowed:", feed.Name)	
	return nil
}
