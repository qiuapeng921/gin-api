package bootstrap

import (
	"context"
	"fmt"
	"gin-api/app/http/middleware"
	"gin-api/app/utility/app"
	"gin-api/app/utility/system"
	"gin-api/app/utility/templates"
	"gin-api/database"
	"gin-api/routers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	app.Error("配置文件", godotenv.Load())

	// 初始Mysql
	app.Panic(app.ConnectDB())
	// 初始Redis
	app.Panic(app.Redis().Connect())
	// 自动创建数据表
	database.AutoGenTable()
	// 汉化参数验证器
	binding.Validator = new(system.Validator)
}

func Run() {
	gin.SetMode(os.Getenv("APP_ENV"))

	engine := gin.Default()

	// 模板初始化
	templates.InitTemplate(engine)
	// 初始化全局中间件
	engine.Use(middleware.RequestLog(), middleware.Cors(), middleware.RequestId())

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
			log.Println("listen:", err)
		}

	}()

	welcome("http://" + endPoint)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.Error("服务关闭", server.Shutdown(ctx))

	log.Println("服务已关闭")
}
