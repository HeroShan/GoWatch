package main

import (
	"math/rand"
	"fmt"
	"time"
	"sync"
)

const (
	numberGoroutines = 4
	taskLoad = 10
)

var (
	wg sync.WaitGroup
	wg2 sync.WaitGroup
	wg3 sync.WaitGroup
)

func init() {
	startime := time.Now()
	fmt.Println(startime)
}

func main() {
	defer func(){
		stoptime := time.Now()
		fmt.Println(stoptime)
	}()
	court := make(chan int)
	wg.Add(3)
	
	go player("Nadal", court)
	go player("Djokovic", court)
	go player("Wkoc", court)
	
	court <- 1
	wg.Wait()
	
	baton := make(chan int)
	wg2.Add(1)
	go Runner(baton)
	
	baton <- 1
	wg2.Wait()
	tasks := make(chan string, taskLoad)
	
	wg3.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}
	
	close(tasks)
	wg3.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	
	for {
		ball, ok := <- court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		
		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			
			close(court)
			return
		}
		
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		
		court <- ball
	}
}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}
	
	time.Sleep(100 * time.Millisecond)
	
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg2.Done()
		return
	}
	
	fmt.Printf("Runner %d Exchange With Runner %d\n", 
		runner,
		newRunner)
	
	baton <- newRunner
}

func worker(tasks chan string, worker int) {
	defer wg3.Done()
	
	for {
		task, ok := <- tasks
		if !ok {
			fmt.Printf("Worker: %d: Shutting Down\n", worker)
			return
		}
		
		fmt.Printf("Worker: %d: Started %s\n", worker, task)
		
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		
		fmt.Printf("Worker: %d: Completed %s\n", worker, task)
	}
}
	
	