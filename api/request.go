package api

import (
	"net/http"

	"database/sql"

	"github.com/labstack/echo"
	"fmt"
)

type Request struct {
	DB *sql.DB
}

func (r *Request) GetTest(c echo.Context) error {

	return c.JSON(http.StatusOK, "GET")
}

func (r *Request) PostTest(c echo.Context) error {
	return c.JSON(http.StatusOK, "POST")
}

func (r *Request) PutTest(c echo.Context) error {
	return c.JSON(http.StatusOK, "PUT")
}

func (r *Request) DeleteTest(c echo.Context) error {
	return c.JSON(http.StatusOK, "DELETE")
}

func (r *Request) Line(c echo.Context) error {
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
	return c.JSON(http.StatusOK, "POST")
}
