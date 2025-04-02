package main

import (
	"github.com/sambakker4/blog_aggregator/internal/config"
	"log"
	"fmt"
	"os"
	"errors"
)

func main() { 
	cfg, err := config.Read()	
	if err != nil {
		log.Fatal(err)
	}
	currentState := state{config: &cfg}
	commandsMap := commands{cmds: make(map[string]func(*state, command)error)}
	commandsMap.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		log.Fatal(errors.New("Must have at least two arguments"))
	}
	currentCommand := command{name: args[1], args: args[2:]}
	commandsMap.run(&currentState, currentCommand)
}
