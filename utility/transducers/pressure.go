package transducers

// 压力传感器
type Pressure struct {
	Tr Transducers
}

func NewPressure(min, max float64, identify, uint string, id int) Pressure {
	return Pressure{
		Tr: NewTransducers(identify, uint,
			NewDataType(FloatType, min, max, EnumTypeStruct{}, BoolTypeStruct{}),
			DefGfFloat(min, max), DefLf, DefSf, id),
	}
}
