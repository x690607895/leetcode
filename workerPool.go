package main

import (
	"fmt"
	"runtime"
)

// 手写协程池

// 任务结构体
type Job struct {
	ID int
}

// 执行结构体
type Worker struct {
	ID         int
	JobChannel chan Job
	Quit       chan struct{}
	Closed     chan struct{}
}

// 池子结构体
type Pool struct {
	WorkNum       int
	JobChannel    chan Job
	WorkerChannel chan Worker
	Quit          chan struct{}
	Closed        chan struct{}
}

// 创建新的worker
func newWorker(id int) Worker {
	return Worker{
		ID:         id,
		JobChannel: make(chan Job),
		Quit:       make(chan struct{}),
		Closed:     make(chan struct{}),
	}
}

// 启动worker
// 将worker塞入池子的workerchannel
// 从jobchannel中取出job，进行job任务
// 从quitchannel中取出quit，关闭jobchannel，向closedchannel中塞入数据，通知外面关闭了
func (w Worker) start(p *Pool) {
	go func() {
		for {
			p.WorkerChannel <- w
			select {
			case job := <-w.JobChannel:
				fmt.Printf("worker Id %d, work %d \n", w.ID, job.ID)
			case <-w.Quit:
				fmt.Printf("worker Id %d, close \n", w.ID)
				close(w.JobChannel)
				w.Closed <- struct{}{}
				return
			}
		}
	}()
}

// 池子启动
// 先根据workNum创建对应数量的worker队列，并且启动
// 如果从jobChannel中读到job，则取出一个workerChannel，并往里面jobChannel塞入一个job
// 如果从quitChannel中读到数据，则循环关闭对应的worker并等待worker的closedchannel返回数据
func (p *Pool) start() {
	for i := 0; i < p.WorkNum; i++ {
		worker := newWorker(i)
		worker.start(p)
	}
	go func() {
		for {
			select {
			case job := <-p.JobChannel:
				worker := <-p.WorkerChannel
				worker.JobChannel <- job
			case <-p.Quit:
				for i := 0; i < p.WorkNum; i++ {
					worker := <-p.WorkerChannel
					worker.Quit <- struct{}{}
					<-worker.Closed
					close(worker.Quit)
					close(worker.Closed)
				}
				p.Closed <- struct{}{}
				return
			}
		}
	}()
}

// 把job塞入池子的job队列
func (p *Pool) addJob(j Job) {
	p.JobChannel <- j
}

// 池子退出，等待池子内所有worker退出，收到了池子的关闭信号，关闭各个池子
func (p *Pool) quit() {
	p.Quit <- struct{}{}
	<-p.Closed
	close(p.JobChannel)
	close(p.WorkerChannel)
	close(p.Quit)
	close(p.Closed)
	runtime.GC()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc / 1024)
	fmt.Println(mem.Sys / 1024)
	fmt.Println("Pool closed")
}

func main() {
	num := 10
	p := Pool{
		WorkNum:       num,
		JobChannel:    make(chan Job),
		WorkerChannel: make(chan Worker, num),
		Quit:          make(chan struct{}),
		Closed:        make(chan struct{}),
	}
	p.start()
	jobNum := 1000000
	for i := 0; i < jobNum; i++ {
		p.addJob(Job{i})
	}
	p.quit()
}
