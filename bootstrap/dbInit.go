package bootstrap

import (
	"gin-api/app/process"
	"gin-api/database"
	"gin-api/helpers/db"
)

func InitTool() {
	db.InitXorm()
	db.InitRedis()
	//queue.InitRabbitMq()
	//db.InitMongo()
	db.InitElastic()
	// 自动创建数据表
	database.AutoGenTable()

	process.InitProcess()
}
