package mywaifulab

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

func SendToSlack() {
	trySend(1)
}

func trySend(trycount int) {
	token := os.Getenv("SLACKTOKEN")
	api := slack.New(token)

	err := api.SetUserPhoto("/tmp/today.png", slack.NewUserSetPhotoParams())
	if err != nil {
		fmt.Println(err)
		if trycount == 1 {
			fmt.Println("Retry")
			trycount++
			trySend(trycount)
		}
		return
	}
	fmt.Println("Success Send")
	return
}
