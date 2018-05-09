# go-reddit-reader
Golang cli client for reading posts/comments/etc from Reddit

go-reddit-reader is a cli tool for interacting with the
Reddit API for fetching posts and comments.

`go-reddit` defaults to fetching the top 10 posts from the
Reddit front page.

`go-reddit -r <subreddit>` fetches top 10 posts from the
specified subreddit.

`go-reddit -p <post-title>` fetch top 10 comments from a
specified post from the front page.

`go-reddit -r <subreddit> -p <post-title>` fetch top 10
comments from a specified post from a specific subreddit.


