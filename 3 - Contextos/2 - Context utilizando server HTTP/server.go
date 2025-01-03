package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request has initialized")
	defer fmt.Println("Request has finalized")

	ctx := r.Context()

	select {

	case <-ctx.Done():
		fmt.Println("Request has canceled")
		return
	case <-time.After(time.Second * 4):
		fmt.Println("Request has processed")
		w.Write([]byte("Request has processed"))
		return
	}
}
