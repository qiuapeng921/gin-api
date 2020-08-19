package bootstrap

import (
	"fmt"
	"gin-api/app/models/admin_role"
	"gin-api/app/models/admins"
	"gin-api/app/models/article_tags"
	"gin-api/app/models/articles"
	"gin-api/app/models/categorys"
	"gin-api/app/models/comments"
	"gin-api/app/models/menus"
	"gin-api/app/models/permissions"
	"gin-api/app/models/role_permission"
	"gin-api/app/models/roles"
	"gin-api/app/models/system_config"
	"gin-api/app/models/tags"
	"gin-api/app/models/user_behavior"
	"gin-api/app/models/users"
	"gin-api/helpers/pool/gredis"
	"gin-api/helpers/pool/grom"
	"gin-api/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	grom.SetUpOrm()
	gredis.SetupRedis()
	autoGenModel()
}

func Run()  {
	gin.SetMode(os.Getenv("APP_ENV"))

	engine := gin.Default()

	gin.ForceConsoleColor()

	// 设置路由
	routers.SetupRouter(engine)

	endPoint := fmt.Sprintf("%s:%s", os.Getenv("HTTP_ADDRESS"), os.Getenv("HTTP_PORT"))

	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf(`sssssssssssssss`)
	log.Printf("welcome use gin,地址:http://127.0.0.1:%s \n", os.Getenv("HTTP_PORT"))
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

// 自动生成数据表
func autoGenModel() {
	client := grom.GetOrm()
	_ = client.Sync2(
		&admins.Entity{},
		&admin_role.Entity{},
		&articles.Entity{},
		&article_tags.Entity{},
		&categorys.Entity{},
		&comments.Entity{},
		&menus.Entity{},
		&permissions.Entity{},
		&roles.Entity{},
		&role_permission.Entity{},
		&system_config.Entity{},
		&tags.Entity{},
		&users.Entity{},
		&user_behavior.Entity{},
	)
}