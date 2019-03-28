package main

import (
	"encoding/json"
	"fmt"
)

func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}

//Json To map
func StringToJson(jsonStr string)(m map[string]map[string]interface{})  {
	err :=json.Unmarshal([]byte(jsonStr),&m)
	checkErr(err)
	fmt.Println(m)
	return
}
//Map To Json
func JsonToString(m map[string]map[string]interface{}) (jsonStr string) {
	bson,err :=json.Marshal(m)
	checkErr(err)
	jsonStr = string(bson)
	fmt.Println(jsonStr)
	return
}
func main() {
	jsonStr := `
		{
			"data": {
				"object":"card",
				"id":"card_123333",
				"first_name":"Hasan",
				"last_name":"URAL",
				"balance":"54.950"
			}
		}
	`
	var m map[string]map[string]interface{}
	m = StringToJson(jsonStr)
	JsonToString(m)



}
