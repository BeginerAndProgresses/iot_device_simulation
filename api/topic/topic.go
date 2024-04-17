package topic

import (
	"github.com/gogf/gf/v2/frame/g"
	"iot_device_simulation/internal/model/entity"
)

type AddReq struct {
	g.Meta           `path:"/" method:"post"`
	PlatForm         string `p:"plat_form" v:"required#平台名不能为空" dc:"平台名"`           // 平台名
	Topic            string `p:"topic" v:"required#topic不能为空" dc:"topic"`           // 通信topic
	FunctionDescribe string `p:"function_describe" v:"required#功能描述不能为空" dc:"功能描述"` // 功能描述
	UserId           int    `p:"user_id" v:"required|integer|min:1#user_id不能为空|user_id只能是整数|最小值不应小于1" dc:"user_id"`
}

type AddRes struct {
	Id int `json:"id" dc:"topic_id"`
}

type EditReq struct {
	g.Meta           `path:"/" method:"put"`
	Id               int    `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
	PlatForm         string `p:"plat_form" v:"required#平台名不能为空" dc:"平台名"`           // 平台名
	Topic            string `p:"topic" v:"required#topic不能为空" dc:"topic"`           // 通信topic
	FunctionDescribe string `p:"function_describe" v:"required#功能描述不能为空" dc:"功能描述"` // 功能描述
}

type EditRes struct {
	Id int `json:"id" dc:"topic_id"`
}

type DelReq struct {
	g.Meta `path:"/" method:"delete"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
}

type DelRes struct {
	Id int `json:"id" dc:"topic_id"`
}

type SearchReq struct {
	g.Meta `path:"/" method:"get"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
}

type SearchRes struct {
	Code  int          `json:"code" dc:"返回码"`
	Topic entity.Topic `json:"topic" dc:"通信topic"`
}

type SearchAllUpTopicReq struct {
	g.Meta `path:"/get_all_up_topic" method:"get"`
	UserId int `p:"user_id" v:"required|integer|min:1#userid不能为空|id只能是整数|最小值不应小于1" dc:"user_id"`
}

type SearchAllUpTopicRes struct {
	Code  int            `json:"code" dc:"是否有返回值"`
	Topic []entity.Topic `json:"topic" dc:"主题"`
}

type SearchAllDownTopicReq struct {
	g.Meta `path:"/get_all_down_topic" method:"get"`
	UserId int `p:"user_id" v:"required|integer|min:1#userid不能为空|id只能是整数|最小值不应小于1" dc:"user_id"`
}

type SearchAllDownTopicRes struct {
	Code  int            `json:"code" dc:"是否有返回值"`
	Topic []entity.Topic `json:"topic" dc:"主题"`
}
