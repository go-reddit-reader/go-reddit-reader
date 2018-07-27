package redditfuncs

import (
	"fmt"

	"github.com/jzelinskie/geddit"
)

// AboutRedditor function
func AboutRedditor(username, password, redditor string) error {
	session, err := geddit.NewLoginSession(
		username,
		password,
		"gedditAgent v1",
	)

	if err != nil {
		panic(err)
	}
	profile, _ := session.AboutRedditor(redditor)
	fmt.Printf("Name: %s\nGold: %t\nMail: %t\n\n", profile.Name, profile.Gold, profile.Mail)
	return nil
}
