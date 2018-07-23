package redditfuncs

import (
	"fmt"

	"github.com/jzelinskie/geddit"
)

// GetSubReddits function
func GetSubReddits(username, password, subreddit string, number int) error {
	session, err := geddit.NewLoginSession(
		username,
		password,
		"gedditAgent v1",
	)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	subOpts := geddit.ListingOptions{
		Limit: number,
	}

	submissions, err := session.SubredditSubmissions(subreddit, geddit.NewSubmissions, subOpts)
	for _, s := range submissions {
		fmt.Printf("Title: %s\nAuthor: %s\n\n", s.Title, s.Author)
	}
	return nil
}
