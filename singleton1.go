package main

import (
	"fmt"
	"log"
	"odd-singleton/singleton"
	"os"
	"strconv"
)

func main() {

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < N; i++ {
		instance := singleton.GetInstance()
		fmt.Printf("%4d  %s\n", i, instance)
	}
}
