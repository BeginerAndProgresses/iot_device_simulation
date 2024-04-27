package transducers

// 烟雾传感器
type Smoke struct {
	Tr Transducers
}

func NewSmoke(identify string, es EnumTypeStruct, id int) Smoke {
	return Smoke{
		Tr: NewTransducers(identify, "", NewDataType(EnumType, 0, 0, es, BoolTypeStruct{}),
			DefGfEnumStruct(es.Enum, es.Describe), DefLf,
			DefSf, id),
	}
}
