package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	data, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", data)
	return nil
}
