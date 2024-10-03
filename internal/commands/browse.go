package commands

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
	"strconv"
)


func HandleBrowse(s *state.GatorState, c Command, user database.User) error {
	var limit int = 2
	if len(c.Args) >= 1 {
		l, err := strconv.Atoi(c.Args[0])
		if err == nil {
			limit = l
		}
	}

	params := database.GetLatestPostsParams {
		Name: user.Name,
		Limit: int32(limit),
	}

	posts, err :=s.Db.GetLatestPosts(context.Background(), params)
	if err != nil {
		return errors.New("Unable to get posts from Db")
	}

	for _, post := range posts {
		fmt.Printf("Title: %v\nDescription: %v\n\n", post.Title, post.Description)
	}
	return nil
}
