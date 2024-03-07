package util

import (
	"encoding/json"
	"fmt"
)

// Transfer
// s1,s2两个结构体结构相同
// 将s1的值传递给s2
// s1和s2的值都应该是指针
func Transfer(s1, s2 interface{}) error {
	marshal, err := json.Marshal(s1)
	if err != nil {
		return fmt.Errorf("Marshal 过程出错 s1:", s1)
	}
	err = json.Unmarshal(marshal, s2)
	if err != nil {
		return fmt.Errorf("Unmarshal 过程出错 marshal:", marshal)
	}
	return nil
}
