# go-reddit-reader
Golang cli client for reading posts/comments/etc from Reddit

go-reddit-reader is a cli tool for interacting with the
Reddit API for fetching posts and comments.

`goreddit` defaults to fetching the top 10 posts from the
Reddit front page.

`goreddit -r <subreddit>` fetches top 10 posts from the
specified subreddit.

`goreddit -p <post-title>` fetch top 10 comments from a
specified post from the front page.

`goreddit -r <subreddit> -p <post-title>` fetch top 10
comments from a specified post from a specific subreddit.


