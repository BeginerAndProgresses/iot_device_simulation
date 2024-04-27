package transducers

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// generatorFunc 生成器函数
type generatorFunc func() interface{}

// linearityFunc 线性度函数
type linearityFunc func(x float64) float64

// sensitivityFunc 灵敏度函数
type sensitivityFunc func(x float64) float64

var (
	DefLf = func(x float64) float64 {
		return x
	}
	DefSf = func(x float64) float64 {
		return 1
	}
	DefGfBoolStruct = func(a, d string) generatorFunc {
		return func() interface{} {
			return BoolTypeStruct{
				Bool:            RandomBool(),
				AbleDescribe:    a,
				DisAbleDescribe: d,
			}
		}
	}
	DefGfEnumStruct = func(enum []interface{}, des map[interface{}]string) generatorFunc {
		return func() interface{} {
			return EnumTypeStruct{
				Enum:     enum,
				data:     RandomEnum(enum),
				Describe: des,
			}
		}
	}

	DefGfInt = func(min, max float64) generatorFunc {
		return func() interface{} {
			return RandomInt(min, max)
		}
	}
	DefGfFloat = func(min, max float64) generatorFunc {
		return func() interface{} {
			return RandomFloat(min, max)
		}
	}
)

type Info struct {
	Data     interface{}
	Identify string
}

// Transducers 传感器接口
type Transducers interface {
	// SendData 每timer秒，向ch发送一次数据
	SendData(timer *time.Ticker, ch chan Info) error
	GetData() interface{}
	Start() error
	Stop() error
	IsOpen() bool
	// Generator 每timer执行一次gf
	Generator(timer *time.Ticker) error
	// SetInterval 设置生成和发送间隔时间
	SetInterval(interval int64)
	GetInterval() int64
	GetId() int
}

// iTransducers 传感器实现
type iTransducers struct {
	state       int // 状态 1 开启 0 关闭
	dt          DataType
	m           sync.Mutex
	gf          generatorFunc
	linearity   linearityFunc
	sensitivity sensitivityFunc
	identify    string // 标识符
	unit        string
	interval    int64
	id          int // 设备id 关联的设备id
}

func (i *iTransducers) Start() error {
	i.m.Lock()
	defer i.m.Unlock()
	i.state = 1
	return nil
}

func (i *iTransducers) Stop() error {
	i.m.Lock()
	defer i.m.Unlock()
	i.state = 0
	return nil
}

func (i *iTransducers) IsOpen() bool {
	return i.state == 1
}

func (i *iTransducers) Generator(timer *time.Ticker) error {
	if !i.IsOpen() {
		return errors.New("设备已经关闭")
	}
	go func() {
		for i.IsOpen() {
			fmt.Println("开始生成数据")
			select {
			case <-timer.C:
				i.m.Lock()
				switch i.dt.types {
				case IntType:
					i.dt.Int = int(i.linearity(float64(i.gf().(int))) * i.sensitivity(1))
				case FloatType:
					i.dt.Float = i.linearity(i.gf().(float64)) * i.sensitivity(1)
				case EnumType:
					i.dt.Enum = i.gf().(EnumTypeStruct)
				case BoolType:
					i.dt.Bool = i.gf().(BoolTypeStruct)
				}
				i.m.Unlock()
				fmt.Println("生成的数据为:", i.dt)
			}
		}
	}()
	return nil
}

func (i *iTransducers) SendData(timer *time.Ticker, ch chan Info) error {
	if !i.IsOpen() {
		return errors.New("设备已经关闭")
	}
	m := sync.Mutex{}
	go func() {
		for i.IsOpen() {
			fmt.Println("开始发送数据")
			select {
			case <-timer.C:
				switch i.dt.types {
				case IntType:
					m.Lock()
					fmt.Println("i.dt.Int", i.dt)
					ch <- Info{Data: i.dt.Int, Identify: i.identify}
					m.Unlock()
				case FloatType:
					m.Lock()
					fmt.Println("i.dt.Float", i.dt)
					ch <- Info{Data: i.dt.Float, Identify: i.identify}
					m.Unlock()
				case EnumType:
					m.Lock()
					fmt.Println("i.dt.Enum", i.dt)
					ch <- Info{Data: i.dt.Enum.data, Identify: i.identify}
					m.Unlock()
				case BoolType:
					m.Lock()
					fmt.Println("i.dt.Bool", i.dt)
					ch <- Info{Data: i.dt.Bool.Bool, Identify: i.identify}
					m.Unlock()
				}
			}
		}
	}()
	return nil
}

func (i *iTransducers) GetId() int {
	return i.id
}

func (i *iTransducers) GetData() interface{} {
	switch i.dt.types {
	case IntType:
		return i.dt.Int
	case FloatType:
		return i.dt.Float
	case EnumType:
		return i.dt.Enum
	case BoolType:
		return i.dt.Bool
	default:
		return struct{}{}
	}
}

func (i *iTransducers) SetInterval(interval int64) {
	i.interval = interval
}

func (i *iTransducers) GetInterval() int64 {
	return i.interval
}

func NewTransducers(identify, unit string, dt DataType, gf generatorFunc, linearity linearityFunc, sensitivity sensitivityFunc, id int) Transducers {
	return &iTransducers{
		identify:    identify,
		state:       1,
		gf:          gf,
		m:           sync.Mutex{},
		linearity:   linearity,
		sensitivity: sensitivity,
		dt:          dt,
		unit:        unit,
		id:          id,
	}
}
