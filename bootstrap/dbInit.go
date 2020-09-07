package bootstrap

import (
	"gin-api/app/process"
	"gin-api/database"
	"gin-api/helpers/db"
	"gin-api/helpers/queue"
	"gin-api/helpers/system"
	"github.com/gin-gonic/gin/binding"
)

func init() {
	binding.Validator = new(system.Validator)
}

func InitTool() {
	db.InitXorm()
	db.InitRedis()
	queue.InitRabbitMq()
	//db.InitMongo()
	db.InitElastic()
	// 自动创建数据表
	database.AutoGenTable()

	process.InitProcess()
}
