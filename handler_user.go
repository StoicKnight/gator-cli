package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	if _, err := s.db.GetUser(context.Background(), name); err != nil {
		log.Fatal("user does not exist:", err)
	}

	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), name)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				log.Fatal("This entry already exists.")
			}
			return fmt.Errorf("could not create new user: %w", err)
		}
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}

	fmt.Printf("User %s has been created\n", user.Name)
	log.Printf("User %s has been created", user.Name)
	return nil
}
