package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var ErrEmptyFile = errors.New("the string is empty")

func CountWords(content string) (int, error) {
	if content == "" {
		return 0, ErrEmptyFile
	} // checking if the string is empty or not

	words := strings.Fields(content)
	// Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.

	return len(words), nil

}

func main() {
	byte, err := os.ReadFile("sample.txt") // reads the file and prints out the content in bytes
	if err != nil {
		log.Println(err)
		return
	}

	// fmt.Println(byte)

	content := string(byte) // converts that byte into the string

	// fmt.Println(content)

	count, err := CountWords(content)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("the number of words in the file : ", count)

	// checking if the file can be deleted

	err = os.Remove("sample.txt")
	if err != nil {
		log.Println(err)
		return
	}
}
