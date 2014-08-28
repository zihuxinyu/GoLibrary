package Library

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)


// map & struct convert is from https://github.com/sdegutis/go.mapstruct
// convert map to struct
func MapToStruct(m map[string]interface{}, s interface{}) {
	v := reflect.Indirect(reflect.ValueOf(s))
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		v.Field(i).Set(reflect.ValueOf(m[key]))
	}
}
// convert struct to map
// s must to be struct, can not be a pointer
func rawStructToMap(s interface{}, snakeCasedKey bool) map[string]interface{} {
	v := reflect.ValueOf(s)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic(fmt.Sprintf("param s must be struct, but got %s", s))
	}
	m := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		if snakeCasedKey {
			key = SnakeCasedName(key)
		}
		val := v.Field(i).Interface()
		m[key] = val
	}
	return m
}
// convert struct to map
func StructToMap(s interface{}) map[string]interface{} {
	return rawStructToMap(s, false)
}
// convert struct to map
// but struct's field name to snake cased map key
func StructToSnakeKeyMap(s interface{}) map[string]interface{} {
	return rawStructToMap(s, true)
}
// get the Struct's name
func StructName(s interface{}) string {
	v := reflect.TypeOf(s)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Name()
}
// load json file to a map
func LoadJsonFile(filePath string) (map[string]interface{}, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	} else if fi.IsDir() {
		return nil, errors.New(filePath + " is not a file.")
	}
	var b []byte
	b, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var conf map[string]interface{}
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
// 获取map的key，返回所有key组成的slice
func MapKeys(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))
	for key, _ := range data {
		keys = append(keys, key)
	}
	return keys
}
// 获取map的key，返回所有key组成的slice
func MapIntKeys(data map[int]int) []int {
	keys := make([]int, 0, len(data))
	for key, _ := range data {
		keys = append(keys, key)
	}
	return keys
}
