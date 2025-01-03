package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writeString("arquivo-writestring")
	readFile("arquivo-writestring")
	removeFile("arquivo-writestring")

	writeBites("arquivo-writebytes")
	readFile("arquivo-writebytes")
	removeFile("arquivo-writebytes")

	readFileBufferIO()
}

func writeString(fileName string) {

	f, err := os.Create(fileName + ".txt")

	panicIfError(err)

	size, err := f.WriteString("Hello, world!")

	panicIfError(err)

	fmt.Printf("\n writestring size %d", size)

	f.Close()
}

func writeBites(fileName string) {

	f, err := os.Create(fileName + ".txt")

	panicIfError(err)

	size, err := f.Write([]byte("Hello, world!"))

	panicIfError(err)

	fmt.Printf("\n writebytes size %d", size)

	f.Close()
}

func removeFile(fileName string) {
	os.Remove(fileName + ".txt")
}

func readFile(fileName string) {

	file, err := os.ReadFile(fileName + ".txt")

	panicIfError(err)

	fmt.Printf("\n Conteúdo do arquivo %s.txt", fileName)
	fmt.Printf("\n %s", string(file))
}

func readFileBufferIO() {

	file, err := os.Open("grande.txt")

	panicIfError(err)

	buffer := make([]byte, 10)
	reader := bufio.NewReader(file)

	fmt.Println("")

	for {

		n, _ := reader.Read(buffer)
		if n == 0 {
			break
		}

		fmt.Printf("\n%d - %s", n, string(buffer[:n]))

	}

	file.Close()
}

func panicIfError(err error) {
	if err != nil {
		panic("Ocorreu algum erro ")
	}
}
