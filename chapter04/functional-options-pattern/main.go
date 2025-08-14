package main

import (
	"log"
	"os"
	"time"

	"golang-programming-essence/chapter04/functional-options-pattern/server"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	svr := server.New(
		"localhost",
		8888,
		server.WithTimeout(time.Minute),
		server.WithLogger(logger),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
