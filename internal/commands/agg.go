package commands

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/internal/feed"
	"gator/internal/state"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func HandlerAgg(s *state.GatorState, c Command) error {
	if len(c.Args) < 1 {
		return errors.New("Must provide a time between requests value")
	}

	duration, err := time.ParseDuration(c.Args[0])
	if err != nil {
		return fmt.Errorf("Not a valid time duration string %v", c.Args[0])
	}

	fmt.Printf("Collecting feeds every %v\n", c.Args[0])

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func scrapeFeeds(s *state.GatorState) error {
	f, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.Db.MarkFeedFetched(context.Background(), f.ID)
	if err != nil {
		return err
	}

	fd, err := feed.FetchFeed(context.Background(), f.Url)
	if err != nil {
		return fmt.Errorf("Error fetching %v: %w", f.Name, err)
	}

	fmt.Println(fd.Channel.Title)
	for _, item := range fd.Channel.Item {
		published, err := convertDate(item.PubDate)
		if err != nil {
			log.Printf("Unknown time format: '%v'\n", item.PubDate)
			return fmt.Errorf("Error parsing time format: %w", err)
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: published,
			FeedID: f.ID,
		}
		_, err = s.Db.CreatePost(context.Background(), params)
		if err != nil {
			if strings.Contains(err.Error(), "unique constraint \"posts_url_key\"") {
				continue
			} else {
				log.Printf("Create Post: %v", err)
			}
		}
	}

	return nil
}

func convertDate(datetime string) (time.Time, error) {
	t, err := time.Parse("Mon, 02 Jan 2006 15:04:05 +0000", datetime)
	if err == nil {
		return t, nil
	}

	return time.Now(), err
}
