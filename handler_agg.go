package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("usage: agg <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("couldn't get next feed to fetch: %v", err)
		return
	}

	if err := s.db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		log.Printf("couldn't mark feed %s as fetched: %v", feed.Name, err)
		return
	}

	fmt.Printf("Fetching %s (%s)...\n", feed.Name, feed.Url)
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't fetch feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf(" * %s\n", item.Title)
	}
	fmt.Printf("Found %d posts in %s\n\n", len(rssFeed.Channel.Item), feed.Name)
}
