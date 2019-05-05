package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Get URL from env
	testUrl := os.Args[1]

	r, err := http.Get(testUrl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.Status)

	if r.StatusCode != http.StatusOK {
		fmt.Println("URL did not respond OK")
	}

}
