package transducers

// 湿度传感器
type Humidity struct {
	Tr Transducers
}

func NewHumidity(min, max float64, identify, uint string, id int) Humidity {
	return Humidity{
		Tr: NewTransducers(identify, uint, DataType{
			Min:   min,
			Max:   max,
			types: FloatType,
		}, DefGfFloat(min, max), DefLf, DefSf, id),
	}
}
