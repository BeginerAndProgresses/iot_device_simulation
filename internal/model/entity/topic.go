// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Topic is the golang structure for table topic.
type Topic struct {
	Id               int    `json:"id"                ` // id
	PlatForm         string `json:"plat_form"         ` // å¹³å°å
	Topic            string `json:"topic"             ` // é€šä¿¡topic
	FunctionDescribe string `json:"function_describe" ` // åŠŸèƒ½æè¿°
	UserId           int    `json:"user_id"           ` // ç”¨æˆ·id
	TType            uint   `json:"t_type"            ` // 订阅为1，上报为0
}
