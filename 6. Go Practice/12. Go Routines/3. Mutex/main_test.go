package main

import (
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"
	var mutext sync.Mutex

	wg.Add(2)
	go updateMsg("x", &mutext)
	go updateMsg("Goodbye, cruel world!", &mutext)
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Error("incorrect value in msg")
	}
}
