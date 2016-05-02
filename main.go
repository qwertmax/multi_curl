package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	} else {
		defer response.Body.Close()
		_, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return 0, err
		}
		return 1, nil
	}
}

func Runner(i int) chan int {
	s := fmt.Sprintf("%d", i)
	status, err := Get("http://www.dmvroll.com/node/" + s)
	if err != nil {
		return <-0
	}
	return <-status
}

func main() {
	id := make(<-chan int)
	done := make(<-chan int)

	for i := 0; i < 100; i++ {
		go Runner(i)
	}

	<-done
}
