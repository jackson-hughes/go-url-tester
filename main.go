package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func testUrl() {

}

func main() {

	if len(os.Args) == 1 {
		log.Fatal("No input provided, please provide at least one url to test")
	}

	numUrls := len(os.Args[1:])

	var wg sync.WaitGroup
	wg.Add(numUrls)

	for _, i := range os.Args[1:] {
		go func(i string) {
			defer wg.Done()
			r, err := http.Get(i)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("url: "+i, "\t"+r.Status)
		}(i)
	}
	wg.Wait()
}
