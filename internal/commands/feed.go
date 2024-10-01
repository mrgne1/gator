package commands

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
	"log"
	"time"

	"github.com/google/uuid"
)

func HandlerAddFeed(s *state.GatorState, c Command) error {
	if len(c.Args) < 1 {
		return errors.New("Must provide a name and url to add a feed")
	} else if len(c.Args) < 2 {
		return ErrNoURL
	}

	user, err := s.Db.GetUser(context.Background(), s.Config.User)
	if err != nil {
		return err
	}

	name := c.Args[0]
	url := c.Args[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feed, err := s.Db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("New Feed:\n%v\n", feed)

	// Add feed follow
	ffparams := database.CreateFeedFollowsParams {
		ID: uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.Db.CreateFeedFollows(context.Background(), ffparams)
	if err != nil {
		return errors.New("Unable to follow")
	}

	fmt.Printf("%v is now following %v\n", user.Name, feed.Name)
	return nil
}

func HandlerFeeds(s *state.GatorState, c Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		log.Println(err)
		return errors.New("Unable to load feeds from DB")
	}

	for _, fd := range feeds {
		user, err := s.Db.GetUserById(context.Background(), fd.UserID)

		var name string
		if err != nil {
			name = "Unknown User"
		} else {
			name = user.Name
		}
		fmt.Printf("%v %v %v\n", fd.Name, fd.Url, name)

	}
	return nil
}

func HandlerFollowing(s *state.GatorState, c Command) error {
	user := s.Config.User

	follows, err := s.Db.GetFeedFollowsForUser(context.Background(), user)
	if err != nil {
		return errors.New("Error finding user feeds")
	}

	for _, f := range follows {
		fmt.Printf("%v\n", f.Feed)
	}

	return nil
}

func HandlerFollow(s *state.GatorState, c Command) error {

	if len(c.Args) < 1 {
		return errors.New("Must provide a URL")
	}

	url := c.Args[0]

	user, err := s.Db.GetUser(context.Background(), s.Config.User)
	if err != nil {
		return err
	}

	feed, err := s.Db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}
	
	ffparams := database.CreateFeedFollowsParams {
		ID: uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.Db.CreateFeedFollows(context.Background(), ffparams)
	if err != nil {
		return errors.New("Unable to follow")
	}

	fmt.Printf("%v is now following %v\n", user.Name, feed.Name)
	return nil
}
