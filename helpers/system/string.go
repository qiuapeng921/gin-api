package system

import (
	"github.com/rs/xid"
	"math/rand"
	"time"
)

// 随机字符串生成
func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成分布式全局唯一id
func GenUUID() string {
	return xid.New().String()
}