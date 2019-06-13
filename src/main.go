package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

func main() {
	token := getenv("SLACKTOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {

			case *slack.MessageEvent:
				info := rtm.GetInfo()

				text := ev.Text
				text = strings.TrimSpace(text)
				text = strings.ToLower(text)

				matched, _ := regexp.MatchString("test", text)

				if ev.User != info.User.ID && matched {
					rtm.SendMessage(rtm.NewOutgoingMessage("_[T]/ Fear the Old Blood _[T]/", ev.Channel))
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
			}
		}
	}
}

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}
