package main

import (
	"fmt"
	"gator/internal/commands"
	"gator/internal/state"
	"os"
	"syscall"
	_ "github.com/lib/pq"
)

func main() {
	state, err := state.NewGatorState()
	if err != nil {
		fmt.Printf("Error initializing state\n%v\n", err)
	}

	cmds := commands.NewCommands()
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerUsers)
	cmds.Register("agg", commands.HandlerAgg)
	cmds.Register("addfeed", commands.MiddlewareLoggedIn(commands.HandlerAddFeed))
	cmds.Register("feeds", commands.HandlerFeeds)
	cmds.Register("following", commands.MiddlewareLoggedIn(commands.HandlerFollowing))
	cmds.Register("follow", commands.MiddlewareLoggedIn(commands.HandlerFollow))
	cmds.Register("unfollow", commands.MiddlewareLoggedIn(commands.HandlerUnfollow))

	if len(os.Args) < 2 {
		fmt.Println("Must provide a command")
		syscall.Exit(1)
	}

	c := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.Run(&state, c)
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}
}
