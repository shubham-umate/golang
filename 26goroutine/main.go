package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup // pointer

var signals = []string{"test"}

var mut sync.Mutex

func main() {
	fmt.Println("Goroutine")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
	}

	for _, v := range websiteList {
		go getStatusCode(v)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)
	//go greeter("Hello")
	//greeter("World")
}

func greeter(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}

}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops in endpoint...")
	}

	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d Status code for %s\n", res.StatusCode, endpoint)

}
