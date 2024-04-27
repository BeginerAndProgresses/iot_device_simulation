package transducers

// 液位传感器
type LiquidLevel struct {
	Tr Transducers
}

func NewLiquidLevel(min, max float64, identify, uint string, id int) LiquidLevel {
	return LiquidLevel{
		Tr: NewTransducers(identify, uint, NewDataType(FloatType, min, max, EnumTypeStruct{}, BoolTypeStruct{}),
			DefGfFloat(min, max), DefLf, DefSf, id),
	}
}
