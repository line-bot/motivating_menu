package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		"18201c5163a1184f1ce066735de7aef9",
		"SQDMB6vhPx3kQGReeVcsYUBzSvupkzzq0u4/jKOX9VTilvz+FEdnm/dgGwogHXmHh7+p4u94t/6PBoR034kgHfYzc5d/pIuMW+vlP8LxIb9KG78QkHHPn5n0EgXx44XIzkbu5lVZSveUYkoiLlTMPgdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("ping\n")
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					fmt.Printf("%v", message)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or
	if err := http.ListenAndServe(":1337", nil); err != nil {
		log.Fatal(err)
	}
}
