package system

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"reflect"
)

// JsonToStruct
func JsonToStruct(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// struct转json
func StructToJson(obj interface{})string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jsonBytes)
}

// json转map
func JsonToMap(jsonStr string) (result map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return result
	}
	return result
}

// map转json
func MapToJson(instance map[string]interface{}) string {
	jsonStr, err := json.Marshal(instance)
	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	return string(jsonStr)
}

// map转struct
func MapToStruct(instance map[string]interface{}, people struct{}) struct{} {
	err := mapstructure.Decode(instance, &people)
	if err != nil {
		fmt.Println(err)
	}
	return people
}

// struct转map
func StructToMap(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		data[objType.Field(i).Name] = objValue.Field(i).Interface()
	}
	return data
}
