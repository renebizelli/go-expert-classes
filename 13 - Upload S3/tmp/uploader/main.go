package main

var {
	s3Client *s3.S3
	s3Bucket string
	wg sync.WaitGroup
}

func init() {

	sess, err := sesion.NewSession(
		&aws.Config{}
	)

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-renebizelli"

}


func main() {

	dir, err := os.Open("./tmp")

	if err != nil {
		panic(err)
	}

	defer dir.Close()

	uploadControl := make(chan struct{}, 100)

	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case filename := <- errorFileUpload:
				uploadControl <- struct{}
				wa.Add(1)
				go uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}

	for {

		files, err := dir.ReadDir(1)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("error reading dir: %s \n", err)
			continue
		}

		wg.Add(1)

		uploadControl <- struct{} // esse ponto segura o For quando atingir o limite de 100, até que algum item do channel seja liberado

		go uploadFile(files[0].Name())
	}

	wg.Wait()

}

func uploadFile(filename string, uploadControl <-chan struct, errorFileUpload chan<- string) {

	defer wg.Done()

	completeFilename := fmt.Sprintf("./tmp/%s", filename)

	fmt.Printf("uploading file %s  to bucket %s \n", completeFilename, s3Bucket)

	f, err := os.Open(completeFilename)

	if err != nil {
		fmt.Printf("error opening file %s\n", completeFilename)
		<- uploadControl
		errorFileUpload <- filename
		return
	}

	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput {
		Bucket : was.String(s3Bucket),
		Key: aws.String(filename),
		Body: f
	})

	if err != nil {
		fmt.Printf("error opening file %s\n", completeFilename)
		<- uploadControl
		errorFileUpload <- filename
		return
	}

	fmt.Printf("file %s uploading successfully\n", completeFilename)

	<- uploadControl
}