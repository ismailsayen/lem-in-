package main

import (
	"fmt"
	"os"

	"lemin/services"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please provide the file name as an argument.")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error: Unable to open the file %s. Please check if the file exists and is accessible.\n", os.Args[1])
		return
	}
	defer file.Close()

	Graph := services.NewGraphData()

	msg := Graph.ValidateFileContent(file)
	if msg != "" {
		fmt.Println(msg)
		return
	}

}
