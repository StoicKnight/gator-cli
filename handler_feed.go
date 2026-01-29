package main

import (
	"context"
	"fmt"

	"github.com/StoicKnight/gator-cli/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %v <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Printf("Feed from '%s' was added to user '%s'\n", feed.Url, s.cfg.CurrentUserName)
	fmt.Println("====================================================================")
	fmt.Println()
	printFeed(feed)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage %v", cmd.Name)
	}
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not list feeds: %w", err)
	}

	fmt.Printf("Found %d feeds\n", len(feeds))
	fmt.Println()
	fmt.Println("================================================================")
	printFeeds(feeds)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func printFeeds(feeds []database.ListFeedsRow) {
	for _, feed := range feeds {
		fmt.Printf("* Name:          %s\n", feed.Name)
		fmt.Printf("* URL:           %s\n", feed.Url)
		fmt.Printf("* UserName:      %s\n", feed.UserName)
		fmt.Println("================================================================")
	}
}
