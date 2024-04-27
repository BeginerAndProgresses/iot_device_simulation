package transducers

// 速度传感器
type Speed struct {
	Tr Transducers
}

func NewSpeed(min, max float64, identify, uint string, id int) Speed {
	return Speed{
		Tr: NewTransducers(identify, uint,
			NewDataType(FloatType, min, max, EnumTypeStruct{}, BoolTypeStruct{}),
			DefGfFloat(min, max),
			DefLf,
			DefSf, id),
	}
}
