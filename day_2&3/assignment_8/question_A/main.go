package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := RemoveFile("a.txt")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("The file has been removed")

}

func RemoveFile(name string) error {
	err := os.Remove(name)
	return err

}
