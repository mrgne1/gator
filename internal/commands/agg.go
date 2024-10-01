package commands

import (
	"fmt"
	"context"
	"gator/internal/state"
	"gator/internal/feed"
)

func HandlerAgg(s *state.GatorState, c Command) error {
	var url string
	if len(c.Args) < 1 {
		//return ErrNoURL
		url = "https://www.wagslane.dev/index.xml"
	} else {
		url = c.Args[0]
	}

	fd, err := feed.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Can't get RSS feed: %v", err)
	}

	fmt.Println(fd)
	return nil
}
