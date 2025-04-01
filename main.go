package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dabates/gator/internal/command"
	"github.com/dabates/gator/internal/config"
	"github.com/dabates/gator/internal/database"
	"github.com/dabates/gator/internal/types"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		// print help
		os.Exit(1)
	}

	state := types.State{
		Config: config.Read(),
	}

	commands := &command.Commands{
		Commands: make(map[string]func(state *types.State, command command.Command) error),
	}

	// setup db
	db, err := sql.Open("postgres", state.Config.DbUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	state.Db = database.New(db)

	setupHandlers(commands)

	err = commands.Run(&state, command.Command{Name: os.Args[1], Args: os.Args[2:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setupHandlers(commands *command.Commands) {
	commands.Register("login", handlerLogin)
	commands.Register("register", handlerRegister)
}

func handlerLogin(s *types.State, cmd command.Command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("invalid arguments: login <username>")
	}

	u, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil || u.Name == "" {
		return fmt.Errorf("user does not exist")
	}

	s.Config.SetUser(u.Name)

	fmt.Println("Logged in", u.Name)

	return nil
}

func handlerRegister(s *types.State, cmd command.Command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("invalid arguments: register <username>")
	}

	u, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err == nil || u.Name != "" {
		return fmt.Errorf("user already exists")
	}

	params := database.CreateUserParams{
		ID:        uuid.NullUUID{UUID: uuid.New(), Valid: true},
		Name:      cmd.Args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	s.Config.SetUser(user.Name)
	fmt.Println("User", user.Name, "created")

	return nil
}
