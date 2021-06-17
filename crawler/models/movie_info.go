package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id                   int64
	Movie_id             int64
	Movie_name           string
	Movie_pic            string
	Movie_director       string
	Movie_writer         string
	Movie_country        string
	Movie_language       string
	Movie_main_character string
	Movie_type           string
	Movie_on_time        string
	Movie_span           string
	Movie_grade          string
}

func init() {
	orm.Debug = true
	_ = orm.RegisterDataBase("default", "mysql", "tester:Root123.@tcp(192.168.146.154:3306)/movies?charset=utf8")
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

func AddMovie(movie_info *MovieInfo) (int64, error) {
	id, err := db.Insert(movie_info)
	return id, err
}

func GetMovieDirector(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	return string(result[0][1])
}

func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span.*?property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	return string(result[0][1])
}

func GetMovieCharacters(movieHtml string) string {
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	//fmt.Println(result)
	mainCharacter := ""
	for _, v := range result {
		mainCharacter += v[1] + "/"
	}
	return strings.Trim(mainCharacter, "/")
}

// <strong class="ll rating_num" property="v:average">6.6</strong>
func GetMovieGrade(movieHtml string) string {
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	//fmt.Println(result)
	return string(result[0][1])
}

// <span property="v:genre">剧情</span>
func GetMovieGenre(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	genre := ""
	for _, v := range result {
		genre += v[1] + "/"
	}
	return strings.Trim(genre, "/")
}

// <span property="v:runtime" content="95">95分钟</span>
func GetMovieOnTime(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	return string(result[0][1])
}

// <span property="v:initialReleaseDate" content="2021-06-12(中国大陆)">2021-06-12(中国大陆)</span>
func GetMovieRunningTime(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	return string(result[0][1])
}

// <a href="https://movie.douban.com/subject/26351864/?from=subject-page">
//                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2500366160.webp" alt="风林火山" class="">
//                </a>
func GetMovieUrls(movieHtml string) []string {
	reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/subject/.*?)"`)
	results := reg.FindAllStringSubmatch(movieHtml, -1)
	//fmt.Println(results)
	//var urls = []string{}
	//return urls
	var movieUrls []string
	for _, v := range results {
		movieUrls = append(movieUrls, v[1])
	}
	return movieUrls
}
