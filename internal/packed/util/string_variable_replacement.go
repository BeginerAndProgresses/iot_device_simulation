package util

import (
	"fmt"
	"strings"
)

// VariableString2String 填充变量字符串，变成完整字符串
func VariableString2String(topic string, args map[string]string) string {
	for k, v := range args {
		vs := fmt.Sprintf("${%s}", k)
		topic = strings.ReplaceAll(topic, vs, v)
	}
	return topic
}
