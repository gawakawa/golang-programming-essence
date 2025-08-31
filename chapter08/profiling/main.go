package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func heavyFunc(wg *sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for range 1000000 {
		s = append(s, "magical pandas")
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		var wg sync.WaitGroup
		wg.Add(1)
		go heavyFunc(&wg)
		wg.Wait()
	}
}
