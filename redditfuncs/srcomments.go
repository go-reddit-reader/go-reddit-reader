package redditfuncs

import (
	"fmt"
	"strconv"
	"time"

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
		s := strconv.FormatFloat(c.Created, 'E', -1, 64)
		t := time.Now()
		created := t.Format(s)
		fmt.Printf("Permalink: %s\nAuthor: %s\nBody: %s\nSubreddit: %s\nUpVotes: %g\nCreated: %s\n\n",
			c.Permalink, c.Author, c.Body, c.Subreddit, c.UpVotes, created)
	}
	if err != nil {
		panic(err)
	}
	return nil
}
