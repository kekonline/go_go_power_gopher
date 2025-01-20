package main

import "sync"

var msg string
var wg sync.WaitGroup

func updateMsg(newMsg string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = newMsg
	m.Unlock()
}

func main() {
	msg = "hello"
	wg.Add(2)
	var mutext sync.Mutex

	go updateMsg("world", &mutext)
	go updateMsg("golang", &mutext)

	wg.Wait()
	println(msg)

}
