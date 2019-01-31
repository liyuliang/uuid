package services

import (
	"github.com/zheng-ji/goSnowFlake"
	"github.com/astaxie/beego"
)

var ids chan int64

var idWorker *goSnowFlake.IdWorker

func Init() {

	if idWorker == nil {
		ids = make(chan int64)
		initIdWorker()
	}

	ids = idChannels()
}

func initIdWorker() {
	wid, _ := beego.AppConfig.Int("workerid")

	if wid <= 0 {
		wid = 1
	}

	iw, err := goSnowFlake.NewIdWorker(int64(wid))
	if err != nil {
		panic(err)
	}
	idWorker = iw
}

func idChannels() chan int64 {

	var out = make(chan int64)

	go func() {
		for {
			out <- <-generator()
		}
	}()

	go func() {
		for {
			out <- <-generator()
		}
	}()

	go func() {
		for {
			out <- <-generator()
		}
	}()

	go func() {
		for {
			out <- <-generator()
		}
	}()

	go func() {
		for {
			out <- <-generator()
		}
	}()

	return out
}

func generator() chan int64 {
	// 创建通道
	out := make(chan int64)
	// 创建协程
	go func() {
		for {
			//向通道内写入数据，如果无人读取会等待

			id, err := getNew()
			if err == nil && id != 0 {
				out <- id
			}
		}
	}()
	return out
}

func getNew() (id int64, err error) {

	id, err = idWorker.NextId()
	return id, err
}

