package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid ticker duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			log.Fatal("Error with scraping feed:", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("could not get next feed: %w", err)
	}

	markedAsFetched, err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("could not mark feed as fetched: %w", err)
	}

	feedData, err := fetchFeed(context.Background(), markedAsFetched.Url)
	if err != nil {
		return fmt.Errorf("count not fetch feed: %w", err)
	}

	fmt.Printf("Titles from '%s' feed (%s):", markedAsFetched.Name, markedAsFetched.Url)
	fmt.Println()
	fmt.Println("==============================================================================")
	for _, item := range feedData.Channel.Item {
		fmt.Printf("* Title: '%s'\n", item.Title)
		fmt.Printf("* Link: '%s'\n", item.Link)
		fmt.Println("==============================================================================")
	}
	return nil
}
