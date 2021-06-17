package routers

import (
	"crawler/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/crawl_movie", &controllers.CrawlMovieController{}, "*:CrawlMovie")
}
