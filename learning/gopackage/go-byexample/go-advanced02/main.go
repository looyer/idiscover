package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type MyContainer struct {
	sync.Mutex
	mapSI map[string]int
}

func (p *MyContainer) inc(name string) {
	p.Lock()
	defer p.Unlock()

	p.mapSI[name]++
}

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("last exception! panic:%v stack:%v", err, string(debug.Stack()))
			os.Exit(-1)
		}
	}()

	log.Printf("~~~~~~~~~ go-advanced02 ~~~~~~~~~")
	log.Printf("now routine num:%v", runtime.NumGoroutine())

	/////////////////////////////////////////////////////////////////
	//Stateful Goroutines
	var readnum, writenum uint64

	readChan := make(chan *readOp)
	writeChan := make(chan *writeOp)

	go func() {
		mapkv := make(map[int]int)
		for {
			select {
			case r := <-readChan:
				r.resp <- mapkv[r.key]
			case w := <-writeChan:
				mapkv[w.key] = w.val
				w.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				r := &readOp{
					key:  rand.Intn(10),
					resp: make(chan int),
				}
				readChan <- r
				<-r.resp
				atomic.AddUint64(&readnum, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 50; i++ {
		go func() {
			for {
				w := &writeOp{
					key:  rand.Intn(10),
					val:  rand.Intn(20),
					resp: make(chan bool),
				}
				writeChan <- w
				<-w.resp
				atomic.AddUint64(&writenum, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	log.Printf("readnum:%v writenum:%v",
		atomic.LoadUint64(&readnum),
		atomic.LoadUint64(&writenum),
	)
	/////////////////////////////////////////////////////////////////

	// /////////////////////////////////////////////////////////////////
	// //Mutexes
	// c := MyContainer{
	// 	mapSI: make(map[string]int),
	// }
	// var wg sync.WaitGroup
	// doInc := func(name string, n int) {
	// 	defer wg.Done()
	// 	for i := 0; i < n; i++ {
	// 		c.inc(name)
	// 	}
	// }
	// wg.Add(4)
	// go doInc("apple", 10000)
	// go doInc("banana", 30000)
	// go doInc("graph", 50000)
	// go doInc("apple", 2000)
	// wg.Wait()
	// log.Printf("mapsi:%v", c.mapSI)
	// /////////////////////////////////////////////////////////////////

	// /////////////////////////////////////////////////////////////////
	// //Atomic Counters
	// var wg sync.WaitGroup
	// var opsafe atomic.Uint32
	// var opun int
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for c := 0; c < 1000; c++ {
	// 			opsafe.Add(1)
	// 			opun++
	// 		}
	// 	}()
	// }
	// wg.Wait()
	// log.Printf("10-goroutine opsafe:%v opun:%v", opsafe.Load(), opun)
	// /////////////////////////////////////////////////////////////////

	// /////////////////////////////////////////////////////////////////
	// //Rate Limiting
	// requests := make(chan int, 5)
	// for i := 0; i < 5; i++ {
	// 	requests <- i
	// }
	// close(requests)
	// limiter := time.Tick(200 * time.Millisecond)
	// for req := range requests {
	// 	<-limiter
	// 	log.Printf("request %v %v", req, time.Now())
	// }

	// burstLimiter := make(chan int, 3)
	// for i := 0; i < 3; i++ {
	// 	burstLimiter <- 0
	// }
	// go func() {
	// 	lim := time.Tick(200 * time.Millisecond)
	// 	for {
	// 		<-lim
	// 		burstLimiter <- 0
	// 	}
	// }()
	// burstyRequests := make(chan int, 5)
	// for i := 0; i < 5; i++ {
	// 	burstyRequests <- (i + 5)
	// }
	// close(burstyRequests)
	// for req := range burstyRequests {
	// 	<-burstLimiter
	// 	log.Printf("burstyRequest %v %v", req, time.Now())
	// }
	// /////////////////////////////////////////////////////////////////
}
