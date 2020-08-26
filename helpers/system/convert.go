package system

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"reflect"
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

//map转json
func MapToJson(instance map[string]interface{}) string {
	jsonStr, err := json.Marshal(instance)
	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	return string(jsonStr)
}

//map转struct
func MapToStruct(instance map[string]interface{}, people struct{}) struct{} {
	err := mapstructure.Decode(instance, &people)
	if err != nil {
		fmt.Println(err)
	}
	return people
}

//struct转map
func StructToMap(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		data[objType.Field(i).Name] = objValue.Field(i).Interface()
	}
	return data
}
