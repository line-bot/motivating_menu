package api

import (
	"net/http"

    "github.com/line/line-bot-sdk-go/linebot"
	"github.com/labstack/echo"
	"fmt"
)

func (r *Request) Test(c echo.Context) error {
    bot, err := linebot.New("18201c5163a1184f1ce066735de7aef9", "SQDMB6vhPx3kQGReeVcsYUBzSvupkzzq0u4/jKOX9VTilvz+FEdnm/dgGwogHXmHh7+p4u94t/6PBoR034kgHfYzc5d/pIuMW+vlP8LxIb9KG78QkHHPn5n0EgXx44XIzkbu5lVZSveUYkoiLlTMPgdB04t89/1O/w1cDnyilFU=")
    if err != nil {

    }
    if _, err := bot.PushMessage("Ubd962e9a51143d8e5240c645b4dacafa", linebot.NewTextMessage("hello")).Do(); err != nil {
        
    }
	return c.JSON(http.StatusOK, "POST")
}
