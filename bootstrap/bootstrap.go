package bootstrap

import (
	"fmt"
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