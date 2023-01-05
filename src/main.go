package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		readFile(s)
	}
	return nil
}

func readFile(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for _, each_ln := range text {
		fmt.Println(each_ln)
	}
}

func main() {
	filepath.WalkDir("/home/santiago/Downloads/enron_mail_20110402", walk)
}
