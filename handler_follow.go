package main

import (
	"context"
	"fmt"

	"github.com/StoicKnight/gator-cli/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not get feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create follow: %w", err)
	}

	fmt.Printf("User '%s' now following feed '%s'\n", s.cfg.CurrentUserName, feed.Name)
	fmt.Println()
	fmt.Println("====================================================================")
	printFeedFollowRow(feedFollow)
	fmt.Println("================================================================")
	return nil
}

func handlerFollowList(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("could not get list of feed follows: %w", err)
	}

	fmt.Printf("User '%s' follows '%d' feeds\n", s.cfg.CurrentUserName, len(follows))
	fmt.Println()
	fmt.Println("================================================================")
	for _, follow := range follows {
		printFeedFollow(follow.UserName, follow.FeedName)
		fmt.Println("================================================================")
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not get feed by url: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not delete feed follow: %w", err)
	}

	fmt.Printf("User %s unfollowed feed %s", user.Name, feed.Name)
	return nil
}

func printFeedFollowRow(follow database.CreateFeedFollowRow) {
	fmt.Printf("* ID:            %s\n", follow.ID)
	fmt.Printf("* Created:       %v\n", follow.CreatedAt)
	fmt.Printf("* Updated:       %v\n", follow.UpdatedAt)
	fmt.Printf("* UserID:        %s\n", follow.UserID)
	fmt.Printf("* UserName:      %s\n", follow.UserName)
	fmt.Printf("* FeedID:        %s\n", follow.FeedID)
	fmt.Printf("* FeedName:      %s\n", follow.FeedName)
}

func printFeedFollow(username string, feedname string) {
	fmt.Printf("* UserName:      %s\n", username)
	fmt.Printf("* FeedName:      %s\n", feedname)
}
