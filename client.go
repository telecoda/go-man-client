package main

import (
	"fmt"
	//"github.com/telecoda/go-man/models"
	"net/http"
	"time"
)

func main() {

	fmt.Println("go-man-client starting")

	start := time.Now()
	//response, err := http.Get("http://localhost:8080/games")
	response, err := http.Post("http://localhost:8080/games", "application/json", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// clear screen
	fmt.Print("\x0c")

	fmt.Println(response.Status)
	// get body
	body := make([]byte, 1076)

	count, err := response.Body.Read(body)

	if err != nil {
		fmt.Println(err)
		return
	}

	if count > 0 {
		fmt.Println(string(body))
	}

	response.Body.Close()

	duration := time.Now().Sub(start)

	fmt.Println(duration)

}
