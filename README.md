# go-reddit-reader
Golang cli client for reading posts/comments/etc from reddit

### Usage

Export your Reddit password locally, `export REDDIT_PASSWORD='password'`

**Main Commands:**

```
NAME:
   redditreader - Returns back information from Reddit

USAGE:
   redditreader [global options] command [command options] [arguments...]

AUTHORS:
   Sudhir Ganesan <ganesansudhir18@gmail.com>
   Luis Ortiz <lortiz1145@gmail.com >

COMMANDS:
     list, l  List posts from subreddit
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

**Sub-Commands:**

```
NAME:
   redditreader list - List posts from subreddit

USAGE:
   redditreader list [command options] [arguments...]

OPTIONS:
   --subreddit value, --sr value  the subreddit
   --sort value                   sort posts [Hot] [New] [Controversial] [Top] [Rising] (default: "hot")
   --number value, -n value       number of posts to return (default: 10)
   --username value, -u value     Reddit username
   --password value, -p value     Reddit password [$REDDIT_PASSWORD]
```

##### Sub-Reddits
**Example 1:**
```
./redditreader list --username='xrod1992' --password=$REDDIT_PASSWORD --subreddit='subaru' --number=15

Title: I love it when we get cool stuff come in on trade-ins
Author: Mr_Stan

Title: 2018 WRX Drone Shot
Author: madn8dawg

Title: I've never been more thrilled with a car decision in my life. About 50 days until she takes me and the kiddo on a cross country adventure.
Author: figureatthegate

Title: Some Wagon Wednesday for you all!!! I still can’t believe I own my dream car
Author: mcjacver

Title: ABS, Traction Control, and Incline lights illuminating on my 2013 Impreza's Dash
Author: gart888

Title: STi Wagon Prodrive conversion
Author: JN_Garage

Title: For Wagon Wednesday, this is my Outback under the stars
Author: darthnut

Title: Does anyone know if they make mesh grille inserts for '18 Crosstrek? Wife's Forester had damge to condenser because a rock got through the holes in the front. If they dont make them, has anyone done a diy job? How would you attach them?
Author: idontliketako

Title: Can anyone tell me what this light is and why its on?
Author: ILoveBlueBalls

Title: A good rear end is hard to beat
Author: eunossti

Title: This Forester may just be a loaner car, but it's still a beautiful Subaru!
Author: komsire22

Title: (Do I know count as a)Wagon Wednesday.
Author: CruelAequitas

Title: Question about short vs long block
Author: IrohtheTeaBender

Title: My friend Crystal Reyes and her BRZ!
Author: CarGirlBecca

Title: ‘14 Crosstrek Hybrid vs ‘18 Crosstrek
Author: dalemg239
```