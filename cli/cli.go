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
					Name:        "sort",
					Usage:       "sort posts [Hot] [New] [Controversial] [Top] [Rising]",
					Value:       "hot",
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
				err := redditfuncs.GetSubReddits(username, password, subreddit, number)
				if err != nil {
					log.Println("Could not return any subreddit submissions")
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
