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
	out := make(chan int)

	s := fmt.Sprintf("%d", i)
	status, err := Get("http://www.dmvroll.com/node/" + s)
	if err != nil {
		out <- 0
		fmt.Printf("nid: %d error\n", i)
		return out
	}

	fmt.Printf("nid: %d\n", i)

	out <- status
	return out
}

func main() {
	// id := make(<-chan int)
	done := make(<-chan int)

	for i := 5000; i <= 50000; i++ {
		fmt.Printf("%d started\n", i)
		go Runner(i)
	}

	<-done
}
