package redditfuncs

import (
	"fmt"

	"github.com/jzelinskie/geddit"
)

// GetSubReddits function
func GetSubReddits(username, password, subreddit, sort string, number int) error {
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

	var sortType geddit.PopularitySort
	switch sort {
	case "new":
		sortType = geddit.NewSubmissions
	case "hot":
		sortType = geddit.HotSubmissions
	case "rising":
		sortType = geddit.RisingSubmissions
	case "top":
		sortType = geddit.TopSubmissions
	case "controversial":
		sortType = geddit.ControversialSubmissions
	default:
		sortType = geddit.DefaultPopularity
	}

	submissions, err := session.SubredditSubmissions(subreddit, sortType, subOpts)
	for _, s := range submissions {
		fmt.Printf("Title: %s\nAuthor: %s\nURL: %s\n\n", s.Title, s.Author, s.URL)
	}
	return nil
}

// SubRedditInfo function
func SubRedditInfo(username, password, subreddit string) error {
	session, err := geddit.NewLoginSession(
		username,
		password,
		"gedditAgent v1",
	)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	info, err := session.AboutSubreddit(subreddit)
	fmt.Printf("%s\n", info)

	return nil
}
