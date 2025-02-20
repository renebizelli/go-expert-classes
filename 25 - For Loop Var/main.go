package main

import "fmt"

func main() {

	// Nova forma de fazer o for i := 0; i < 10; i++ {}
	for i := range 10 {
		fmt.Println(i)
	}

	//bug For Loop Var

	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v) // na versão antiga, essa variavel V era a mesma em todas as go-routines, sempre o ultimo valor do array.
			done <- true
		}()
	}

	for range values {
		<-done
	}

}
