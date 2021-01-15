package main

import (
	"fmt"
	"net/http"
	"sync"

	u "github.com/rumbletumjum/ping/utils"
)

func main() {
	file := "urls.txt"
	urls, err := u.ReadLines(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", file)
	}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error pinging: %v\n", url)
			}
			fmt.Printf("Successful ping: %v\n", url)
		}(url)
	}
	wg.Wait()
}
