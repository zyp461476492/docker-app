package database

import (
	"errors"
	"github.com/asdine/storm"
	"log"
	"sync"
)

var ErrPoolClosed = errors.New("资源池已经关闭")

type Pool struct {
	lock      sync.Mutex
	resources chan *storm.DB
	factory   func() (*storm.DB, error)
	closed    bool
}

func (p *Pool) Acquire() (*storm.DB, error) {
	select {
	case r, flag := <-p.resources:
		log.Printf("从通道 %s 中请求资源", r.Bolt.GoString())
		var err error = nil
		if !flag {
			err = ErrPoolClosed
		}
		return r, err
	default:
		log.Print("无可用资源，创建新资源")
		return p.factory()
	}
}

func (p *Pool) Release(db *storm.DB) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 资源池如果关闭，关闭掉现有的连接
	if p.closed {
		err := db.Close()
		if err != nil {
			log.Printf("关闭连接时发生异常,原因 %s", err.Error())
		}
		return
	}
	select {
	case p.resources <- db:
		log.Printf("释放连接 %s 至资源池", db.Bolt.GoString())
	default:
		// 资源池已满，关闭当前的连接
		log.Printf("释放连接 %s 至资源池失败，当前资源池已满，转为关闭连接", db.Bolt.GoString())
		err := db.Close()
		if err != nil {
			log.Printf("关闭连接时发生异常,原因 %s", err.Error())
		}
	}
}

func (p *Pool) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.closed = true
	// 先关闭通道，随后关闭连接
	close(p.resources)
	for r := range p.resources {
		err := r.Close()
		if err != nil {
			log.Printf("关闭连接时发生异常,原因 %s", err.Error())
		}
	}
}
