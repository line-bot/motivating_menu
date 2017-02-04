package phaseOne

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

// テスト方法
// goapp serve ./app
// curl http://localhost:8080/api/phase1 で　下のようなjsonが帰ってきたらおｋ

// 下みたいなstringで返せばパイソンさんが頑張ってくれます！
// お願いします！がんばってください！

/*
フォーマット例
[
    {
        "category": "1",
        "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/hamburger_240.jpeg"
    },
    {
        "category": "2",
        "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/pizza_240.jpeg"
    },
    {
        "category": "3",
        "https://s3-us-west-2.amazonaws.com/lineapitest/bread_240.jpeg"
    }
]
*/

// 実際のデータ
var res = `[ { "category": "1", "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/hamburger_240.jpeg" }, { "category": "2", "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/pizza_240.jpeg" }, { "category": "3", "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/bread_240.jpeg" } ]`

type PhaseOne struct {
	DB *sql.DB
}

func (p *PhaseOne) Response(c echo.Context) error {
	// 乱数で送信するカテゴリ画像とカテゴリNoを決定する

	// Jsonフォーマットに沿ったstringを作成する
	// ※構造体でいい感じに出来るなら、それでもおｋ。出来上がればなんでも！

	return c.String(http.StatusOK, res)
}
