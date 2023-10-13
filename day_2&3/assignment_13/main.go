package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var ErrEmptyFile = errors.New("the file is empty")

func CreateFile(name string) (*os.File, error) {
	f, err := os.Create(name)
	if err != nil {
		return f, err
	}
	return f, nil
}

func ReadOpenedFile(f *os.File) ([]byte, error) {
	store := make([]byte, 100)
	_, err := f.Read(store)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func ReadFile(path string) ([]byte, error) {
	byte, err := os.ReadFile(path) // reads the file and prints out the content in bytes
	if err != nil {
		return nil, err
	}
	return byte, nil
}

func CountWords(store []byte) (int, error) {
	content := string(store) // converts that byte into the string

	if content == "" {
		return 0, ErrEmptyFile
	} // checking if the string is empty or not

	words := strings.Fields(content)
	return len(words), nil

}

func main() {
	f, err := CreateFile("sample1.txt")
	if err != nil {
		log.Println("Could not ceated file ; ", err)
		return
	}
	store, err := ReadOpenedFile(f)
	if err != nil {
		log.Println("Could not read the file ; ", err)
		return
	}
	len, err := CountWords(store)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("The number of words in the file : ", len)
}
