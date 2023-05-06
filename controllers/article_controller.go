package controllers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"

	"fmt"
	"github.com/labstack/echo/v4"
	"go-api-sample/models"
	"log"
	"os"
	"strconv"
)

type ArticleController struct{}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

func (ac *ArticleController) GetArticles(c echo.Context) error {
	var articles []models.Article

	dsn := getConString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}

	page, _ := strconv.Atoi(c.Param("page"))

	pageCount, err := strconv.Atoi(c.Param("pageCount"))
	if err != nil {
		pageCount = 10
	}

	db.Limit(pageCount).Offset(pageCount * (page - 1)).Find(&articles)

	return c.JSON(http.StatusOK, articles)
}

func getConString() string {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
}
