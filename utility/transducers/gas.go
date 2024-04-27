package transducers

// Gas 气体传感器
type Gas struct {
	Tr Transducers
	// 气体传感器的参数
	id int // 设备id
}

func NewGas(id, min, max int, identify, uint string) Gas {
	return Gas{
		Tr: NewTransducers(identify, uint,
			NewDataType(IntType, float64(min), float64(max), EnumTypeStruct{}, BoolTypeStruct{}),
			DefGfInt(float64(min), float64(max)),
			DefLf, DefSf, id),
	}
}
