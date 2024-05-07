package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/gosimple/slug"
)

type News struct {
	Title	string `json:"title"`
	Url		string `json:"url"`
	BaseUrl string `json:"baseurl"`
}

var GroupNews []News
var baseUrl = "/"

func main(){
	c := colly.NewCollector()
	c.OnHTML("li.day-article", func(e *colly.HTMLElement) {
		news := News {
			Title: e.ChildText("a"),
			Url: baseUrl + slug.Make(strings.ToLower(e.ChildText("a"))),
			BaseUrl: e.ChildAttr("a", "href"),
		}
		GroupNews = append(GroupNews, news)
	})

	router := gin.Default()
	router.GET("/:year/:month", func(ctx *gin.Context){
		Year := ctx.Param("year")
		Month := ctx.Param("month")
		c.Visit("https://www.pcgamer.com/archive/"+ Year + "/"+ Month +"/")
		ctx.JSON(http.StatusOK, gin.H{"message": "success","count": len(GroupNews),"data": GroupNews})
	})
	router.Run(":3000")
}