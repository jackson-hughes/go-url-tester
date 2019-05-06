package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("No input provided, please provide at least one url to test")
	}

	for _, i := range os.Args[1:] {

		r, err := http.Get(i)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("url: "+i, "\t"+r.Status)
	}
}
