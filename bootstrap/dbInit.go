package bootstrap

import (
	"gin-api/database"
	"gin-api/helpers/db"
	"gin-api/helpers/queue"
)

func Init()  {
	db.InitXorm()
	db.InitRedis()
	db.InitMongo()
	db.InitElastic()
	queue.InitRabbitMq()
	// 自动创建数据表
	database.AutoGenTable()

}
