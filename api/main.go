package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	// going backwards after 0, 1 pair is not supported
	if (f.previous != 0) {
		f.current, f.previous = f.previous, f.current - f.previous
	}
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
	http.HandleFunc("/previous", previousHandler)
	http.HandleFunc("/next", nextHandler)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	fibLock.Lock()
	current := fib.current
	fibLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(map[string]string {
		"current": fmt.Sprintf("%v", current),
	})

	if err != nil {
		log.Fatalf("Error while marshaling to json. Err: %s", err)
	}
	
	w.Write(resp)
	return

}

func previousHandler(w http.ResponseWriter, r *http.Request) {
	fibLock.Lock()
	fib.bwd()
	current := fib.current
	fibLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(map[string]string {
		"current": fmt.Sprintf("%v", current),
	})
	
	if err != nil {
		log.Fatalf("Error while marshaling to json. Err: %s", err)
	}

	w.Write(resp)
	return

}

func nextHandler(w http.ResponseWriter, r *http.Request) {
	fibLock.Lock()
	fib.fwd()
	current := fib.current
	fibLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(map[string]string {
		"current": fmt.Sprintf("%v", current),
	})

	if err != nil {
		log.Fatalf("Error while marshaling to json. Err: %s", err)
	}

	w.Write(resp)
	return

}