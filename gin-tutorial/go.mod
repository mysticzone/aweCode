module gin-tutorial

go 1.14

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/go-ini/ini v1.62.0
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
)

//replace (
//		gin-tutorial/pkg/settings => ~/src/gin-tutorial/pkg/settings
//		gin-tutorial/conf    	  => ~/src/gin-tutorial/pkg/conf
//		gin-tutorial/middleware  => ~/src/gin-tutorial/middleware
//		gin-tutorial/models 	  => ~/src/gin-tutorial/models
//		gin-tutorial/routers 	  => ~/src/gin-tutorial/routers
//)
