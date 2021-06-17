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
		fmt.Println("GetQeueueLength: ", length)
		if length == 0 {
			break
		}
		sUrl = models.PopFromQueue()
		fmt.Println("sUrl: ", sUrl)
		if models.IsVisit(sUrl) {
			continue
		}
		//resp := httplib.Get(sUrl)
		resp := httplib.Get(sUrl)
		sMovieHtml, err := resp.String()
		if err != nil {
			fmt.Printf("CrawlMovie Error: %v\n", err)
			panic(err)
		}

		movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
		fmt.Println("Movie_name: ", movieInfo.Movie_name)
		if movieInfo.Movie_name != "" {
			movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character = fmt.Sprintf("%v", models.GetMovieCharacters(sMovieHtml))
			movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_span = models.GetMovieOnTime(sMovieHtml)
			movieInfo.Movie_type = models.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_country = models.GetMovieRunningTime(sMovieHtml)
			movieInfo.Movie_on_time = "2021-06-12"
			fmt.Println("===================================")
			fmt.Println("MovieInfo: ", movieInfo)
			fmt.Println("===================================")
			m_id, err := models.AddMovie(&movieInfo)
			fmt.Printf("MovieID: %v\t Err:%v\n", m_id, err)
		}

		urls := models.GetMovieUrls(sMovieHtml)
		fmt.Printf("sMovieHtml Urls: %v\n", urls)
		for _, url := range urls {
			models.PutInQueue(url)
			c.Ctx.WriteString("<br>" + url + "</br>")
		}

		models.AddToSet(sUrl)
		time.Sleep(time.Second)
	}

}
