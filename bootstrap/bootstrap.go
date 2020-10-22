package bootstrap

import (
	"context"
	"fmt"
	"gin-api/app/process"
	"gin-api/app/utility/app"
	"gin-api/app/utility/db"
	"gin-api/app/utility/queue"
	"gin-api/app/utility/system"
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
	err := godotenv.Load()
	if err != nil {
		panic("配置文件不存在" + err.Error())
	}
	// 初始化所有工具类
	db.InitXorm()
	db.InitRedis()
	system.SecurePanic(app.Redis().Connect())

	queue.InitRabbitMq()
	//db.InitMongo()
	db.InitElastic()
	// 自动创建数据表
	database.AutoGenTable()

	process.InitProcess()
	
	// 汉化参数验证器
	binding.Validator = new(system.Validator)
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

	welcome("http://" + endPoint)

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
