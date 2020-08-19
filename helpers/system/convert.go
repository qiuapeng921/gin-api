package system

import (
	"encoding/json"
	"fmt"
	"time"
)

func JsonToStruct(jsonStr string, Struct struct{}) struct{} {
	err := json.Unmarshal([]byte(jsonStr), &Struct)
	if err != nil {
		fmt.Println(err)
	}
	return Struct
}

//struct转json
func StructToJson(Struct struct{}) {
	jsonBytes, err := json.Marshal(Struct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}

//json转map
func JsonToMap(jsonStr string) (result map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return result
	}
	return result
}

// 时间戳转时间格式
func FormatDate(timestamp int64) string {
	layout := "2006-01-02 15:04:05"
	return time.Unix(timestamp, 0).Format(layout)
}