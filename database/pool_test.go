package database

import (
	"github.com/zyp461476492/docker-app/types"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

//定义全局常量
const (
	maxGoroutines = 20 //使用25个goroutine模拟同时的连接请求
	poolSize      = 2  //资源池中的大小
)

func getDb() (interface{}, error) {
	config := types.Config{
		FileLocation: "test.db",
		Timeout:      10,
	}
	return GetStorm(config)
}

func TestGetStorm(t *testing.T) {
	config := types.Config{
		FileLocation: "test.db",
		Timeout:      0,
	}
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func(gid int) {
			db, err := GetStorm(config)
			//defer CloseStorm(db)
			if err != nil {
				log.Printf("新增时发生异常 %s ", err.Error())
			} else {
				log.Printf("资源信息：%d %s", gid, db)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	// 同时并发数量
	wg.Add(maxGoroutines)
	myPool, err := NewStormPool(1, 1, 2, getDb)
	if err != nil {
		t.Errorf("获取连接池失败，原因：%s", err.Error())
	}
	for i := 0; i < maxGoroutines; i++ {
		//模拟请求
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		go func(gid int) {
			execQuery(gid, myPool)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//定义一个查询方法,参数是当前 gorotine Id和资源池
func execQuery(goroutineId int, pool *StormPool) {
	//从池里请求资源,第一次肯定是没有的,就会创建一个dbConn实例
	conn, _ := pool.Acquire()
	//睡眠一下,模拟查询过程
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	log.Printf("执行查询...协程ID [%d] 资源ID [%s]", goroutineId, conn)
	//将创建的dbConn实例放入了资源池的缓冲通道里
	defer pool.Release(conn)
}
