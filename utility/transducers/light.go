package transducers

// 光传感器
type Light struct {
	Tr Transducers
}

func NewLight(min, max float64, identify, uint string, id int) Light {
	return Light{
		Tr: NewTransducers(identify, uint,
			NewDataType(IntType, min, max, EnumTypeStruct{}, BoolTypeStruct{}),
			DefGfInt(min, max), DefLf, DefSf, id),
	}
}
