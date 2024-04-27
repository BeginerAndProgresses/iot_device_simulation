package transducers

import (
	"math/rand"
	"time"
)

const (
	IntType   = "int"
	FloatType = "float"
	BoolType  = "bool"
	EnumType  = "enum"
)

type BoolTypeStruct struct {
	Bool            int
	AbleDescribe    string
	DisAbleDescribe string
}

type EnumTypeStruct struct {
	Enum     []interface{}
	data     interface{}
	Describe map[interface{}]string
}

type DataType struct {
	types string
	Enum  EnumTypeStruct // 如果是枚举值，存值
	Bool  BoolTypeStruct // 布尔值
	Int   int
	Float float64
	Min   float64 // 其他类型最大最小值
	Max   float64
}

func NewDataType(types string, min, max float64, enums EnumTypeStruct, bs BoolTypeStruct) DataType {
	if types == BoolType {
		return DataType{types: types, Bool: bs}
	}
	if types == IntType || types == FloatType {
		return DataType{
			types: types,
			Min:   min,
			Max:   max,
		}
	}
	if types == EnumType {
		return DataType{
			types: types,
			Enum:  enums,
		}
	}
	return DataType{
		types: types,
	}
}

func RandomInt(min, max float64) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(int(max-min)) + int(min)
}

func RandomFloat(min, max float64) float64 {
	rand.Seed(time.Now().Unix())
	return rand.Float64()*(max-min) + min
}

func RandomBool() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(2) % 2
}

func RandomEnum(enum []interface{}) interface{} {
	rand.Seed(time.Now().Unix())
	return enum[rand.Intn(len(enum))]
}
