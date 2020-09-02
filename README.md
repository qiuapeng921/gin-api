# gin-api

## 安装
```
git clone https://gitea.phpswoole.cn/qiuapeng921/gin-api $GOPATH/src/gin-api
```

## 如何运行

### 必须

- Mysql
- Redis
- ElasticSearch
- RabbitMq

### 运行
```
cd $GOPATH/src/gin-api
cp .env.example .env

###修改.env配置
DB_HOST=127.0.0.1:3306
DB_DATABASE=test
DB_USERNAME=root
DB_PASSWORD=
DB_CHARSET=utf8mb4
DB_MAX_IDLE=20
DB_MAX_OPEN=100

REDIS_HOST=127.0.0.1:6379
REDIS_PASSWORD=
REDIS_MAX_IDLE=50
REDIS_MAX_ACTIVE=1200

go run main.go
```

## 特性
- Gin
- Xorm
- Redis
- Mongo
- ElasticSearch
- RabbitMq
- Socket
- Templates