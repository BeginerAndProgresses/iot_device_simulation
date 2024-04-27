package transducers

type TransducerBoolOption struct {
	AbleDescription    string `json:"able_description" dc:"为true描述"`
	DisAbleDescription string `json:"disable_description" dc:"为false描述"`
}

type TransducerIntegerOption struct {
	Min  int    `json:"min" dc:"最小值"`
	Max  int    `json:"max" dc:"最大值"`
	Unit string `json:"unit" dc:"单位"`
}

type TransducerFloatOption struct {
	Min  float64 `json:"min" dc:"最小值"`
	Max  float64 `json:"max" dc:"最大值"`
	Unit string  `json:"unit" dc:"单位"`
}

type TransducerEnumOption struct {
	Enum map[interface{}]string `json:"enum" dc:"枚举"`
}
