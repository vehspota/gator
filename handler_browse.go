package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/vehspota/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		n, err := strconv.Atoi(cmd.args[0])
		if err != nil || n < 1 {
			return fmt.Errorf("limit must be a positive number, got %q", cmd.args[0])
		}
		limit = n
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts yet. Run 'agg' to collect some.")
		return nil
	}

	for _, post := range posts {
		date := "unknown date"
		if post.PublishedAt.Valid {
			date = post.PublishedAt.Time.Format("Mon Jan 2 2006")
		}
		fmt.Printf("%s from %s\n", date, post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		if post.Description.Valid {
			fmt.Printf("    %.200s\n", post.Description.String)
		}
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
