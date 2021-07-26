package system

import (
	"encoding/json"
	"github.com/goinggo/mapstructure"
	"reflect"
)

func JsonToStruct(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return err
	}
	return nil
}

func StructToJson(obj interface{}) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func JsonToMap(jsonStr string) (result map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return result
	}
	return result
}

func MapToJson(instance map[string]interface{}) string {
	jsonStr, err := json.Marshal(instance)
	if err != nil {
		panic(err)
	}
	return string(jsonStr)
}

func MapToStruct(instance map[string]interface{}, people struct{}) struct{} {
	err := mapstructure.Decode(instance, &people)
	if err != nil {
		panic(err)
	}
	return people
}

func StructToMap(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		data[objType.Field(i).Name] = objValue.Field(i).Interface()
	}
	return data
}
