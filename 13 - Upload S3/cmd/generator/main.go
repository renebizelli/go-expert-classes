package main

func main() {

	i := 0 
	
	for {
		f, err := os.Create(fmt.Sprintf(".tmp/file%d.txt", i++))
		if err != nil {
			panic(err)
		}

		defer f.Close()

		f.WriteString("Hello, World")

		fmt.Println(i)
	}

}