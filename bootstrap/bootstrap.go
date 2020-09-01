package bootstrap

import (
	"context"
	"fmt"
	"gin-api/database"
	"gin-api/helpers/db"
	"gin-api/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	db.SetUpXorm()
	db.SetupRedis()

	// 自动创建数据表
	database.AutoGenTable()
}

func Run() {
	gin.SetMode(os.Getenv("APP_ENV"))

	engine := gin.Default()

	gin.ForceConsoleColor()

	// 设置路由
	routers.SetupRouter(engine)

	address := os.Getenv("HTTP_ADDRESS")
	port := os.Getenv("HTTP_PORT")
	endPoint := fmt.Sprintf("%s:%s", address, port)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}()

	welcome("http://"+endPoint)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("服务关闭错误:", err.Error())
	}
	log.Println("服务已关闭")
}

const logo = `
 _____ _          ___        _ 
|  __ (_)        / _ \      (_)
| |  \/_ _ __   / /_\ \_ __  _ 
| | __| | '_ \  |  _  | '_ \| |
| |_\ \ | | | | | | | | |_) | |
 \____/_|_| |_| \_| |_/ .__/|_|
                      | |      
                      |_|      `

func welcome(endPoint string) {
	fmt.Println(strings.Replace(logo, "*", "`", -1))
	fmt.Println("")
	fmt.Println(fmt.Sprintf("Server      Name:     %s", os.Getenv("APP_NAME")))
	fmt.Println(fmt.Sprintf("System      Name:     %s", runtime.GOOS))
	fmt.Println(fmt.Sprintf("Go          Version:  %s", runtime.Version()[2:]))
	fmt.Println(fmt.Sprintf("Gin         Version:  %s", gin.Version))
	fmt.Println(fmt.Sprintf("Listen      Address:  %s", endPoint))
}