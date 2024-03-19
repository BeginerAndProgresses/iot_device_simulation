package util

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ChangeJsonPostId 修改Json字符串中的id
func ChangeJsonPostId(platform, oleJson string, newId int) (newJson string) {
	if platform == "华为云" {
		return oleJson
	}
	mapjson := make(map[string]interface{})
	err := json.Unmarshal(String2Bytes(oleJson), &mapjson)
	if err != nil {
		fmt.Println("转换失败mapjson")
	}
	switch platform {
	case "阿里云":
		mapjson["id"] = strconv.Itoa(newId)
	case "腾讯云":
		mapjson["clientToken"] = strconv.Itoa(newId)

	default:

	}
	marshal, err := json.Marshal(mapjson)
	if err != nil {
		fmt.Println("marshals", marshal)
	}
	return Bytes2String(marshal)
}
