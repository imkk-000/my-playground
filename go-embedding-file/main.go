package main

import (
	"embed"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//go:embed hello.txt
	//go:embed html/* image/*
	var content embed.FS
	log.Println(content.ReadFile("hello.txt"))
	dirs, err := content.ReadDir("html")
	failOnError(err, "fail to read dir")
	for _, dir := range dirs {
		info, err := dir.Info()
		failOnError(err, "fail to get file info")
		log.Printf("print info: %+v", info)
	}
	dirs, err = content.ReadDir("image")
	failOnError(err, "fail to read dir")
	for _, dir := range dirs {
		info, err := dir.Info()
		failOnError(err, "fail to get file info")
		log.Printf("print info: %+v", info)
	}

	//go:embed mock_main.go
	var contentFile string
	log.Println(contentFile)
}
