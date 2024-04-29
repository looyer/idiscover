package main

import (
	"log"
	"runtime"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("~~~~~~~~~ go-advanced ~~~~~~~~~")
	log.Printf("now routine num:%v", runtime.NumGoroutine())

	// ///////////////////////////////////////////////////////////////////////////
	// //WaitGroups
	// worker := func(id int) {
	// 	log.Printf("worker :%v starting...", id)
	// 	time.Sleep(time.Second)
	// 	log.Printf("worker :%v finish.", id)
	// }
	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(id int) {
	// 		defer wg.Done()
	// 		worker(id)
	// 	}(i)
	// }
	// wg.Wait()
	// ///////////////////////////////////////////////////////////////////////////

	// ///////////////////////////////////////////////////////////////////////////
	// //Worker Pools
	// timeStart := time.Now()
	// worker := func(id int, jobs <-chan int, results chan<- int) {
	// 	for j := range jobs { //receive values from the channel repeatedly it is closed.
	// 		log.Printf("worker:%v handle job:%v", id, j)
	// 		time.Sleep(time.Second)
	// 		results <- id
	// 		log.Printf("worker:%v finish job:%v", id, j)
	// 	}
	// }
	// const numJobs = 5
	// chJobs := make(chan int, numJobs)
	// chResults := make(chan int, numJobs)
	// for i := 0; i < 3; i++ {
	// 	go worker(i, chJobs, chResults)
	// }
	// for i := 1; i <= numJobs; i++ {
	// 	chJobs <- i
	// }
	// close(chJobs)
	// for i := 0; i < numJobs; i++ {
	// 	<-chResults
	// }
	// mills := time.Since(timeStart).Milliseconds()
	// log.Printf("use real time:%.4fs", float64(mills)/1000)
	// ///////////////////////////////////////////////////////////////////////////

	// ///////////////////////////////////////////////////////////////////////////
	// //Tickers do repeatedly at regular intervals
	// ticker := time.NewTicker(500 * time.Millisecond)
	// done := make(chan bool)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case t := <-ticker.C:
	// 			log.Printf("tick at %v", t)
	// 		}
	// 	}
	// }()
	// time.Sleep(1600 * time.Millisecond)
	// ticker.Stop()
	// done <- true
	// ///////////////////////////////////////////////////////////////////////////

	// ///////////////////////////////////////////////////////////////////////////
	// //Timers do once in the future!
	// timer1 := time.NewTimer(2 * time.Second)
	// <-timer1.C
	// log.Printf("timer1 fire one")

	// timer2 := time.NewTimer(time.Second)
	// go func() {
	// 	<-timer2.C //blocking there! //
	// 	log.Printf("timer 2 fired")
	// }()
	// stop := timer2.Stop()
	// if stop {
	// 	log.Printf("timer2 stopped")
	// }
	// time.Sleep(3 * time.Second)
	// ///////////////////////////////////////////////////////////////////////////

	// ///////////////////////////////////////////////////////////////////////////
	// queue := make(chan string, 3)
	// queue <- "apple"
	// queue <- "pear"
	// queue <- "banana"
	// close(queue)
	// for fruit := range queue {
	// 	log.Printf("grap fruit is %v", fruit)
	// }
	// ///////////////////////////////////////////////////////////////////////////

	// ///////////////////////////////////////////////////////////////////////////
	// jobs := make(chan int, 5)
	// done := make(chan bool)
	// go func() {
	// 	for {
	// 		j, ok := <-jobs
	// 		if ok {
	// 			log.Printf("do work: %v", j)
	// 		} else {
	// 			log.Printf("work finish.")
	// 			done <- true
	// 			return
	// 		}
	// 	}
	// }()

	// for i := 0; i < 3; i++ {
	// 	jobs <- (i + 1)
	// }
	// close(jobs)
	// <-done
	// ///////////////////////////////////////////////////////////////////////////

	// ///////////////////////////////////////////////////////////////////////////
	// c1 := make(chan string, 1)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c1 <- "late res01"
	// }()
	// select {
	// case res := <-c1:
	// 	log.Printf("get external res:%v", res)
	// case <-time.After(3 * time.Second):
	// 	log.Printf("external timeout!")
	// }
	// ///////////////////////////////////////////////////////////////////////////

	// ///////////////////////////////////////////////////////////////////////////
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func(s string) {
	// 	defer wg.Done()
	// 	log.Print(s)
	// 	log.Printf("now routine num:%v", runtime.NumGoroutine())
	// }("another routine")

	// chanmsg := make(chan string)
	// go func() {
	// 	defer func() { <-chanmsg }()
	// 	log.Print("use chan sync.")
	// 	log.Printf("now routine num:%v", runtime.NumGoroutine())
	// }()

	// chanmsg <- "ok"
	// wg.Wait()
	// ///////////////////////////////////////////////////////////////////////////

	// ///////////////////////////////////////////////////////////////////////////
	// chaninput := make(chan string, 4)
	// go holdPrint(chaninput)

	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	word := input.Text()
	// 	if word == "q" || word == "Q" {
	// 		break
	// 	}
	// 	chaninput <- word
	// }
	// ///////////////////////////////////////////////////////////////////////////

	log.Printf("now routine num:%v", runtime.NumGoroutine())
}

func holdPrint(chaninput <-chan string) {
	for {
		word := <-chaninput
		log.Printf("recv your input: %v", word)
	}
}
