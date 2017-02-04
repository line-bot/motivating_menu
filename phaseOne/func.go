package phaseOne

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

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
// var res = `[ { "category": "1", "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/hamburger_240.jpeg" }, { "category": "2", "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/pizza_240.jpeg" }, { "category": "3", "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/bread_240.jpeg" } ]`

type PhaseOne struct {
	DB *sql.DB
}

// imageは画像URLとcategoryを持ちます
type Image struct {
	category string
	ImageURL string
}

func createImages() []Image {

	image1 := Image{
		category: "1",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619284/d3d3681a-eb34-11e6-9f74-f513ec33182a.jpeg",
	}

	image2 := Image{
		category: "2",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619296/135cc1a2-eb35-11e6-83ff-884fbfdb00e8.jpg",
	}

	image3 := Image{
		category: "3",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619306/546cefd2-eb35-11e6-8679-dac4ea32744c.jpg",
	}

	image4 := Image{
		category: "4",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619310/78fd2bb4-eb35-11e6-9ab0-6e3028262131.jpg",
	}

	image5 := Image{
		category: "5",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619318/b36b904c-eb35-11e6-8167-8e6e28c817df.jpeg",
	}

	image6 := Image{
		category: "6",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619326/d91039ba-eb35-11e6-98e5-182e416c50fd.jpg",
	}

	image7 := Image{
		category: "7",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619334/14edaae4-eb36-11e6-9453-1330495134af.jpeg",
	}

	image8 := Image{
		category: "8",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619339/3e38e558-eb36-11e6-8778-4776e13cf313.jpg",
	}

	image9 := Image{
		category: "9",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg",
	}

	image10 := Image{
		category: "10",
		ImageURL: "https://cloud.githubusercontent.com/assets/25542931/22619393/2d648cc2-eb37-11e6-80be-1f53691533ad.jpeg",
	}

	var images []Image
	images = append(images, image1, image2, image3, image4, image5, image6, image7, image8, image9, image10)
	return images
}

func (p *PhaseOne) Response(c echo.Context) error {
	// 乱数で送信するカテゴリ画像とカテゴリNoを決定する

	rand.Seed(time.Now().Unix())
	images := createImages()
	selectedImage1 := images[rand.Intn(len(images))]
	selectedImage2 := images[rand.Intn(len(images))]
	selectedImage3 := images[rand.Intn(len(images))]

	var selectedImages []Image
	selectedImages = append(selectedImages, selectedImage1, selectedImage2, selectedImage3)

	// Jsonフォーマットに沿ったstringを作成する
	// ※構造体でいい感じに出来るなら、それでもおｋ。出来上がればなんでも！
	imagesJSON, err := json.Marshal(selectedImages)
	if err != nil {
		log.Println(err)
		return err
	}

	var buf bytes.Buffer
	buf.Write(imagesJSON)

	return c.JSON(http.StatusOK, buf.String())
	//return nil, imagesJSON
}
