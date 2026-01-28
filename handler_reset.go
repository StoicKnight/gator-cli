package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		log.Fatal("Could not delete users")
	}

	fmt.Println("Successfully reset DB")
	return nil
}
