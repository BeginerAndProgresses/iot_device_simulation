package transducers

// 温度传感器
type Temperature struct {
	Tr Transducers
}

func NewTemperature(min, max float64, identify, uint string, id int) Temperature {
	return Temperature{
		Tr: NewTransducers(
			identify,
			uint,
			NewDataType(
				FloatType,
				min,
				max,
				EnumTypeStruct{},
				BoolTypeStruct{},
			),
			DefGfFloat(min, max),
			DefLf,
			DefSf, id)}
}
