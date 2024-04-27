package transducers

// 红外传感器
type Infrared struct {
	Tr Transducers
}

func NewInfrared(identify string, bs BoolTypeStruct, id int) Infrared {
	return Infrared{
		Tr: NewTransducers(identify, "",
			NewDataType(BoolType, 0, 0, EnumTypeStruct{}, bs),
			DefGfBoolStruct(bs.AbleDescribe, bs.DisAbleDescribe), DefLf, DefSf,
			id),
	}
}
