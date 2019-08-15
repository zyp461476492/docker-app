package database

import (
	"errors"
	"github.com/asdine/storm"
	"github.com/zyp461476492/docker-app/types"
	"log"
	"sync"
	"time"
)

var (
	ErrInvalidConfig = errors.New("invalid pool config")
	ErrPoolClosed    = errors.New("pool closed")
)

type factory func() (interface{}, error)

type Pool interface {
	Acquire() (interface{}, error)
	Release(interface{}) error // 释放资源
	Close(interface{}) error   // 关闭资源
	Shutdown() error           // 关闭池
}

type StormPool struct {
	sync.Mutex
	pool        chan interface{}
	maxOpen     int  // 池中最大资源数
	numOpen     int  // 当前池中资源数
	minOpen     int  // 池中最少资源数
	closed      bool // 池是否已关闭
	maxLifetime time.Duration
	factory     factory
}

func NewStormPool(minOpen, maxOpen int, maxLifetime time.Duration, factory factory) (*StormPool, error) {
	if maxOpen <= 0 || minOpen > maxOpen {
		return nil, ErrInvalidConfig
	}
	p := &StormPool{
		maxOpen:     maxOpen,
		minOpen:     minOpen,
		maxLifetime: maxLifetime,
		factory:     factory,
		pool:        make(chan interface{}, maxOpen),
	}

	config := types.Config{
		FileLocation: "test.db",
		Timeout:      10,
	}

	var wg sync.WaitGroup
	wg.Add(minOpen)
	for i := 0; i < minOpen; i++ {
		go func(gid int) {
			resource, err := GetStorm(config)
			if err != nil {
				log.Printf("新建资源池时，新建资源失败，%s", err.Error())
			} else {
				p.numOpen++
				p.pool <- resource
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	return p, nil
}

func (p *StormPool) Acquire() (interface{}, error) {
	if p.closed {
		return nil, ErrPoolClosed
	}
	for {
		resource, err := p.getOrCreate()
		if err != nil {
			return nil, err
		}
		// todo maxLifttime 最大等待时间
		return resource, err
	}
}

func (p *StormPool) getOrCreate() (interface{}, error) {
	select {
	case resource := <-p.pool:
		return resource, nil
	default:
	}
	p.Lock()
	if p.numOpen >= p.maxOpen {
		resource := <-p.pool
		p.Unlock()
		return resource, nil
	}
	// 新建连接
	resource, err := p.factory()
	if err != nil {
		p.Unlock()
		return nil, err
	}
	p.numOpen++
	p.Unlock()
	return resource, nil
}

// 释放单个资源到连接池
func (p *StormPool) Release(resource interface{}) error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	p.pool <- resource
	p.Unlock()
	return nil
}

// 关闭单个资源
func (p *StormPool) Close(resource interface{}) error {
	p.Lock()
	db := resource.(*storm.DB)
	CloseStorm(db)
	p.numOpen--
	p.Unlock()
	return nil
}

// 关闭连接池，释放所有资源
func (p *StormPool) Shutdown() error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	close(p.pool)
	for resource := range p.pool {
		db := resource.(*storm.DB)
		CloseStorm(db)
		p.numOpen--
	}
	p.closed = true
	p.Unlock()
	return nil
}
