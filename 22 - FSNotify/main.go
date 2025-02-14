package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var dbconfig DBConfig

func marshallConfig(file string) {

	data, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &dbconfig)

	if err != nil {
		panic(err)
	}

}

func main() {

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	done := make(chan bool)
	go func() {

		for {

			select {

			case event, ok := <-watcher.Events: // fica observando se houve alguma mudança no arquivo
				if !ok {
					return
				}

				fmt.Println("Caso aconteça algum evento: event:", event)

				if event.Op&fsnotify.Write == fsnotify.Write { // evento de escrita
					fmt.Println("modified file:", event.Name)
					marshallConfig("config.json")
					fmt.Println(dbconfig)
				}

			case err, ok := <-watcher.Errors: // caso aconteça algum erro

				if !ok {
					return
				}

				fmt.Println("error:", err)
			}

		}

	}()

	err = watcher.Add("config.json")

	if err != nil {
		panic(err)
	}

	marshallConfig("config.json")

	fmt.Println(dbconfig)

	<-done

}
