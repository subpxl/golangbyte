package main

import (
	"fmt"
	"os"
	"strconv"
)

var topics = []string{
	"speeding up with grpc",
	"deployment with kubernetes",
	"up and running with aws",
	"go microservies frameworks",
	"security best practices for microservices",
	"use of pprof tool",
	"use of sync package",
	"use of garbage collection",
	"use fo generics",
}

func createMainFile(folderName string) error {
	filePath := fmt.Sprintf("%s/main.go", folderName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func main() {
	for index, topic := range topics {
		folderName := strconv.Itoa(index+1) + " " + topic
		err := os.MkdirAll(folderName, 0755)
		if err != nil {
			fmt.Printf("Error creating folder %s: %s\n", topic, err)
			continue
		}

		err = createMainFile(folderName)
		if err != nil {
			fmt.Printf("Error creating main.go in %s: %s\n", topic, err)
		}
	}
}
