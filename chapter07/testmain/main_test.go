package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Before")
	ret := m.Run()
	log.Println("After")
	os.Exit(ret)
}

func TestA(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
