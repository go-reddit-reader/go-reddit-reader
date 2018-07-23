package redditfuncs

import (
	"fmt"

	"github.com/jzelinskie/geddit"
)

// GetSubReddits function
func GetSubReddits(username, password, subreddit string, number int) error {
	session, _ := geddit.NewLoginSession(
		username,
		password,
		"gedditAgent v1",
	)

	subOpts := geddit.ListingOptions{
		Limit: number,
	}

	submissions, _ := session.SubredditSubmissions(subreddit, geddit.NewSubmissions, subOpts)
	for _, s := range submissions {
		fmt.Printf("Title: %s\nAuthor: %s\n\n", s.Title, s.Author)
	}
	return nil
}
