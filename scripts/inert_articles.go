package main

import (
	"encoding/json"
	"go-api-sample/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
)

func main() {
	// NewsAPI の APIにリクエストして chatCPTに関すニュースを取得する
	url := "https://newsapi.org/v2/everything?q=chatGPT&from=2023-05-03&sortBy=publishedAt&apiKey=f1871965ddc848cbaac48e37c52a10fa"

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	// JSONを構造体にエンコード
	var response map[string][]map[string]string
	json.Unmarshal(body, &response)

	dsn := "root:password@tcp(127.0.0.1:3306)/news_article?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	for _, v := range response["articles"] {
		article := models.Article{
			Title:        v["title"],
			Description:  v["description"],
			Content:      v["content"],
			ArticleUrl:   v["url"],
			ImageUrl:     v["urlToImage"],
			ResourceName: v["author"],
		}

		db.Create(&article)
	}
}
