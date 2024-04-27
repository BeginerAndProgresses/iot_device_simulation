package transducers

import (
	"fmt"
	"sync"
	"time"
)

var (
	CTransducers *transducersController
)

func init() {
	CTransducers = NewTransducersController()
	CTransducers.Start()
}

type transducersController struct {
	Transducerss   map[Transducers]bool
	m              sync.RWMutex
	RegisterChan   chan Transducers
	UnRegisterChan chan Transducers
	CloseChan      chan Transducers
	Infos          chan Info
}

func NewTransducersController() *transducersController {
	return &transducersController{
		m:              sync.RWMutex{},
		Transducerss:   make(map[Transducers]bool),
		RegisterChan:   make(chan Transducers, 100),
		UnRegisterChan: make(chan Transducers, 100),
		CloseChan:      make(chan Transducers, 100),
		Infos:          make(chan Info, 100),
	}
}

func (c *transducersController) Register(t Transducers) {
	c.m.Lock()
	defer c.m.Unlock()
	c.Transducerss[t] = true
	c.RegisterChan <- t
}

func (c *transducersController) UnRegister(t Transducers) {
	fmt.Println("UnRegister")
	// 死锁，导致请求阻塞
	//c.m.Lock()
	//defer c.m.Unlock()
	c.Transducerss[t] = false
	c.UnRegisterChan <- t
}

func (c *transducersController) Close(t Transducers) {
	c.m.Lock()
	defer c.m.Unlock()
	delete(c.Transducerss, t)
	c.CloseChan <- t
}

func (c *transducersController) RangeAll(fc func(t Transducers, value bool) bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	for transducers, b := range c.Transducerss {
		if !fc(transducers, b) {
			return
		}
	}
}

func (c *transducersController) Start() {
	go func() {
		for {
			select {
			case tr := <-c.RegisterChan:
				_ = tr.Start()
				ticker := time.NewTicker(time.Duration(tr.GetInterval()*2/3) * time.Second)
				_ = tr.Generator(ticker)
				_ = tr.SendData(ticker, c.Infos)
			case tr := <-c.UnRegisterChan:
				_ = tr.Stop()
			case <-c.CloseChan:
				// 不知道要干啥，先不写
			}
		}

	}()
}
