package fuzzlib

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func WordListParser(filename string) ([]string, error) {
	wordlist := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}

	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		word := Scanner.Text()
		trimmedWord := strings.Trim(word, "\f\t\r\n ")

		wordlist = append(wordlist, trimmedWord)
	}

	return wordlist, nil
}
