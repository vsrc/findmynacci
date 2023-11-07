package main

import (
	"fmt"
	"net/http"
	"sync"
)

// data structure to hold the fibonacci numbers
type findmynacci struct {
	current uint64
	previous uint64
}

// function to initialize the findmynacci struct
func newFindmynacci() *findmynacci {
	return &findmynacci{current: 1, previous: 0}
}

// steps forward with the fibonacci sequence
func (f *findmynacci) fwd() {
	f.current, f.previous = f.current + f.previous, f.current
}

// steps backward with the fibonacci sequence
func (f *findmynacci) bwd() {
	f.current, f.previous = f.previous, f.current - f.previous
}

var (
	// global variable to hold the findmynacci struct
	fib = newFindmynacci()

	// global variable to lock and prevent the race conditions
	fibLock sync.Mutex
)

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/current", currentHandler)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	fibLock.Lock()
	current := fib.current
	fibLock.Unlock()

	fmt.Fprintln(w, current)
}