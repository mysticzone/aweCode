package controllers

import (
	"crawler/models"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {
	var movieInfo models.MovieInfo
	// connect redis
	models.ConnectRedis("192.168.146.54:6379")

	sUrl := "https://movie.douban.com/subject/29984000/"
	models.PutInQueue(sUrl)

	for {
		length := models.GetQueueLength()
		if length == 0 {
			break
		}
		sUrl = models.PopFromQueue()
		if models.IsVisit(sUrl) {
			continue
		}

		resp := httplib.Get(sUrl)
		sMovieHtml, err := resp.String()
		if err != nil {
			panic(err)
		}

		movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
		if movieInfo.Movie_name != "" {
			movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character = fmt.Sprintf("%v", models.GetMovieCharacters(sMovieHtml))
			movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_span = models.GetMovieOnTime(sMovieHtml)
			movieInfo.Movie_type = models.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_country = models.GetMovieRunningTime(sMovieHtml)
			movieInfo.Movie_on_time = "2021-06-12"

			m_id, err := models.AddMovie(&movieInfo)
			fmt.Println(m_id, err)
		}

		urls := models.GetMovieUrls(sMovieHtml)
		for _, url := range urls {
			models.PutInQueue(url)
			c.Ctx.WriteString("<br>" + url + "</br>")
		}

		models.AddToSet(sUrl)
		time.Sleep(time.Second)

	}

	//movieId, err := models.AddMovie(&movieInfo)
	//fmt.Printf("MovieID:%v\tErr:%v\n", movieId, err)
	//
	//c.Ctx.WriteString(models.GetMovieDirector(sMovieHtml) + "|")
	//c.Ctx.WriteString(models.GetMovieName(sMovieHtml) + "|")
	//c.Ctx.WriteString(fmt.Sprintf("%v", models.GetMovieCharacters(sMovieHtml)) + "|")
	//c.Ctx.WriteString(models.GetMovieGrade(sMovieHtml) + "|")
	//c.Ctx.WriteString(models.GetMovieGenre(sMovieHtml) + "|")
	//c.Ctx.WriteString(models.GetMovieOnTime(sMovieHtml) + "|")
	//c.Ctx.WriteString(models.GetMovieRunningTime(sMovieHtml) + "|")

	//
	//fmt.Printf("Urls:%v\n", urls)
}
