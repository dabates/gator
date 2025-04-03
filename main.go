package main

import (
	"database/sql"
	"fmt"
	"github.com/dabates/gator/internal/command"
	"github.com/dabates/gator/internal/config"
	"github.com/dabates/gator/internal/database"
	"github.com/dabates/gator/internal/types"
	_ "github.com/lib/pq"
	"os"
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
	commands.Register("login", command.LoginHandler)
	commands.Register("register", command.RegisterHandler)
	commands.Register("reset", command.ResetHandler)
	commands.Register("users", command.UsersHandler)
	commands.Register("agg", command.AggHandler)
	commands.Register("addfeed", middlewareLoggedIn(command.AddFeedHandler))
	commands.Register("feeds", command.FeedsHandler)
	commands.Register("follow", middlewareLoggedIn(command.FollowHandler))
	commands.Register("following", middlewareLoggedIn(command.FollowingHandler))
	commands.Register("unfollow", middlewareLoggedIn(command.UnfollowHandler))
}
