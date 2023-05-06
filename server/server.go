package server

import (
	"github.com/labstack/echo/v4"
	"go-api-sample/controllers"
)

type Server struct {}

// routing の定義を追加する
func ServerStart() {
	e := echo.New()
	ac := controllers.NewArticleController()

	e.GET("/articles/:page", ac.GetArticles)

	e.Logger.Fatal(e.Start(":8080"))
}
