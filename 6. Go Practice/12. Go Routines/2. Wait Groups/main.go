package main

import (
	"fmt"
	"sync"
)

func main() {
	words := []string{"Hello", "World", "!", "love", "you", "!", "code", "in", "go"}

	// import should use the pointer of waitgroup
	var wg sync.WaitGroup
	wg.Add(len(words))

	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}

	wg.Wait()

}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)

}
