package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job1 struct {
	id       int
	randomno int
}
//type Result1 struct {
//	job         Job1
//	sumofdigits int
//}

var noOfWorkers = 10

var jobs1 = make(chan Job1, noOfWorkers)
//var results1 = make(chan Result1, 10)

//func digits1(number int) int {
//	sum := 0
//	no := number
//	for no != 0 {
//		digit := no % 10
//		sum += digit
//		no /= 10
//	}
//	time.Sleep(2 * time.Second)
//	return sum
//}
func worker1(wg *sync.WaitGroup) {
	for job := range jobs1 {
		//output := Result1{job, digits1(job.randomno)}
		////results <- output
		//fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", output.job.id, output.job.randomno, output.sumofdigits)
		fmt.Printf("Job id %d, input random no %d\n", job.id, job.randomno)
		//fmt.Println(job)
	}
	wg.Done()
}
func createWorkerPool1(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		fmt.Println("start worker ", i)
		wg.Add(1)
		go worker1(&wg)
	}
	wg.Wait()
	//close(results)
}
func allocate1(noOfJobs int) {
	//fmt.Println("sleep 15s, then create jobs")
	//for i := 0; i < 30; i++ {
	//	fmt.Println("sleep ", i)
	//	time.Sleep(500 * time.Millisecond)
	//}
	//fmt.Println("create jobs")
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job1{i, randomno}
		jobs1 <- job
		//fmt.Println(len(jobs))
	}
	close(jobs1)
}

//func result1(done chan bool) {
//	for result := range results1 {
//		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
//	}
//	done <- true
//}
func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate1(noOfJobs)
	//done := make(chan bool)
	//go result(done)
	//noOfWorkers := 10
	createWorkerPool1(noOfWorkers)
	//<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
