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
	"syscall"
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

	startLogo()
	fmt.Println("访问地址:" + "http://" + endPoint)
	fmt.Println("进程Id:", syscall.Getpid())
	fmt.Println("goVersion:", runtime.Version())
	fmt.Println("ginVersion:", gin.Version)

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

func startLogo() {
	logo := `
 _____ _          ___        _ 
|  __ (_)        / _ \      (_)
| |  \/_ _ __   / /_\ \_ __  _ 
| | __| | '_ \  |  _  | '_ \| |
| |_\ \ | | | | | | | | |_) | |
 \____/_|_| |_| \_| |_/ .__/|_|
                      | |      
                      |_|      `
	fmt.Println(logo)
}
