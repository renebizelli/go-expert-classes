package main

func main() {

	canal := make(chan string)

	go func() {
		canal <- "Hello"
	}()

	mensagem := <-canal
	println(mensagem)

}
