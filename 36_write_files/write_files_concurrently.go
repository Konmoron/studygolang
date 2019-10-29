package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consumer(data chan int, done chan bool) {
	f, err := os.Create("./36_write_files/concurrent.txt")
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}

	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i:=0; i < 999; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consumer(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <- done
	if d == true {
		fmt.Println("file write successfully")
	} else {
		fmt.Println("file write failed")
	}
}
