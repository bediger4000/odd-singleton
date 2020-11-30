package main

import (
	"fmt"
	"log"
	"odd-singleton/singleton"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func main() {

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(num int) {
			instance := singleton.GetInstance()
			fmt.Printf("%4d  %s\n", num, instance)
			wg.Done()
		}(i)
		runtime.Gosched()
	}
	wg.Wait()
}
