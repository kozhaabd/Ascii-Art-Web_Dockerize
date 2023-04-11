package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func Getstrings(filename string) ([]string, error) {
	var text []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error in filename!!!")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		text = append(text, str)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return text, nil
}
