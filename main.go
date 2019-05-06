package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	for _, i := range os.Args[1:] {

		fmt.Println(i) // for debugging

		r, err := http.Get(i)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(r.Status)
	}
}
