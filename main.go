package main

import (
"context"

"cloud.google.com/go/storage"

"io"
"log"
"os"

"google.golang.org/api/option"
)

func main() {
	credentialFilePath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	// クライアントを作成する
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	// GCSオブジェクトを書き込むファイルの作成
	f, err := os.Open("iamge/sample.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// オブジェクトのReaderを作成
	bucketName := "sample-go"
	objectPath := "icon2"

	w := client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	defer w.Close()

	if _, err := io.Copy(w, f); err != nil {
		log.Println(err)
	}
	if err := w.Close(); err != nil {
		log.Println(err)
	}
	log.Println("done")
}

