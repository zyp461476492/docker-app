package database

import (
	"errors"
	"github.com/asdine/storm"
	"log"
	"sync"
)

var ErrPoolClosed = errors.New("资源池已经关闭")

type StormPool struct {
	lock      sync.Mutex
	resources chan *storm.DB
	factory   func() (*storm.DB, error)
	closed    bool
}

func NewPool(fn func() (*storm.DB, error), size uint) *StormPool {
	pool := StormPool{
		resources: make(chan *storm.DB, size),
		factory:   fn,
	}
	return &pool
}

func (p *StormPool) Acquire() (*storm.DB, error) {
	select {
	case r, flag := <-p.resources:
		log.Printf("从通道 %s 中请求资源", r)
		var err error = nil
		if !flag {
			err = ErrPoolClosed
		}
		return r, err
	default:
		log.Print("无可用资源，创建新资源")
		db, err := p.factory()
		if err != nil {
			log.Printf("获取资源发生异常 %s", err.Error())
		} else {
			log.Printf("获取资源成功 %s", db)
		}
		return db, err
	}
}

func (p *StormPool) Release(db *storm.DB) {
	p.lock.Lock()
	log.Printf("开始释放资源 %s", db)
	defer p.lock.Unlock()
	// 资源池如果关闭，关闭掉现有的连接
	if p.closed {
		CloseStorm(db)
		return
	}
	if db == nil {
		log.Printf("连接资源为空,无法进行释放！")
	} else {
		select {
		case p.resources <- db:
			log.Printf("释放连接 %s 至资源池", db)
		default:
			// 资源池已满，关闭当前的连接
			log.Printf("释放连接 %s 至资源池失败，当前资源池已满，转为关闭连接", db)
			CloseStorm(db)
		}
	}
}

func (p *StormPool) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.closed = true
	// 先关闭通道，随后关闭连接
	close(p.resources)
	for r := range p.resources {
		CloseStorm(r)
	}
}
