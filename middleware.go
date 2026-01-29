package main

import (
	"context"
	"fmt"

	"github.com/StoicKnight/gator-cli/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("could not get user: %w", err)
		}

		return handler(s, c, user)
	}
}
