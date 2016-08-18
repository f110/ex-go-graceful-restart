package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				res, err := http.Get("http://127.0.0.1:8000")
				if err != nil {
					log.Fatal(err)
				}
				defer res.Body.Close()

				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				}
				if string(body) != "Hello world!" {
					log.Fatal("Can't get content")
				}
			}
		}()
	}
	wg.Wait()
}
