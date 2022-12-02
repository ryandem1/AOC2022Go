package common

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// ReadLinesFromFile will read an input file located in the inputs directory and will send the string lines through a
// channel
func ReadLinesFromFile(fileName string) chan string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(filepath.Join(cwd, "inputs", fileName))
	f, err := os.Open(filepath.Join(cwd, "inputs", fileName))

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	lines := make(chan string)
	go func() {
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
