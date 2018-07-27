package cli

import (
	"log"
	"os"

	"github.com/go-reddit-reader/redditreader/redditfuncs"
	"github.com/urfave/cli"
)

// Cli function that executes the cli code
func Cli() {
	app := cli.NewApp()
	app.Name = "redditreader"
	app.Usage = "Returns back information from Reddit"
	app.HideVersion = true
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Sudhir Ganesan",
			Email: "ganesansudhir18@gmail.com",
		},
		cli.Author{
			Name:  "Luis Ortiz",
			Email: "lortiz1145@gmail.com ",
		},
	}
	var subreddit string
	var sort string
	var number int
	var username string
	var password string
	var reddituser string
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List posts from subreddit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "subreddit, sr",
					Usage:       "the subreddit",
					Destination: &subreddit,
				},
				cli.StringFlag{
					Name:  "sort",
					Usage: "sort posts [hot] [new] [controversial] [top] [rising]",
					//Value:       "hot",
					Destination: &sort,
				},
				cli.IntFlag{
					Name:        "number, n",
					Usage:       "number of posts to return",
					Value:       10,
					Destination: &number,
				},
				cli.StringFlag{
					Name:        "username, u",
					Usage:       "Reddit username",
					Destination: &username,
				},
				cli.StringFlag{
					Name:        "password, p",
					Usage:       "Reddit password",
					EnvVar:      "REDDIT_PASSWORD",
					Destination: &password,
				},
			},
			Action: func(c *cli.Context) error {
				//redditfuncs.GetSubReddits(username, password, subreddit, sort, number)
				redditfuncs.GetSubReddits(username, password, subreddit, sort, number)
				return nil
			},
		},
		{
			Name:    "comments",
			Aliases: []string{"c"},
			Usage:   "List comments from subreddit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "subreddit, sr",
					Usage:       "the subreddit",
					Destination: &subreddit,
				},
				cli.StringFlag{
					Name:        "username, u",
					Usage:       "Reddit username",
					Destination: &username,
				},
				cli.StringFlag{
					Name:        "password, p",
					Usage:       "Reddit password",
					EnvVar:      "REDDIT_PASSWORD",
					Destination: &password,
				},
			},
			Action: func(c *cli.Context) error {
				redditfuncs.SubRedditComments(username, password, subreddit)
				return nil
			},
		},
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "Info about a user or subreddit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "subreddit, sr",
					Usage:       "the subreddit",
					Destination: &subreddit,
				},
				cli.StringFlag{
					Name:        "username, u",
					Usage:       "Your Reddit username",
					Destination: &username,
				},
				cli.StringFlag{
					Name:        "password, p",
					Usage:       "Your Reddit password",
					EnvVar:      "REDDIT_PASSWORD",
					Destination: &password,
				},
				cli.StringFlag{
					Name:        "reddituser, ru",
					Usage:       "The Reddit username you want get info on",
					Destination: &reddituser,
				},
			},
			Action: func(c *cli.Context) error {
				if len(reddituser) > 0 {
					redditfuncs.AboutRedditor(username, password, reddituser)
				} else if len(subreddit) > 0 {
					redditfuncs.SubRedditInfo(username, password, subreddit)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
