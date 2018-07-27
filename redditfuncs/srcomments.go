package redditfuncs

import (
	"fmt"

	"github.com/jzelinskie/geddit"
)

// SubRedditComments function
func SubRedditComments(username, password, subreddit string) error {
	session, err := geddit.NewLoginSession(
		username,
		password,
		"gedditAgent v1",
	)

	comments, err := session.SubredditComments(subreddit)
	for _, c := range comments {
		fmt.Printf("Comment: %s\n", c)
	}
	if err != nil {
		panic(err)
	}
	return nil
}
