package fuzzlib

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Config struct {
	Threads   int
	Logger    Logger
	Processor Processor
	Wordlist  []string
}

type safeCounter struct {
	mutex   sync.Mutex
	counter int
}

func NewConfig(threads int, wordlist []string, processor Processor) Config {
	logger := NewLogger()
	return Config{
		Threads:   threads,
		Logger:    logger,
		Processor: processor,
		Wordlist:  wordlist,
	}
}

func (c Config) Run() {
	wordChan := make(chan string, c.Threads)
	resultChan := make(chan string, c.Threads)
	errorChan := make(chan string, c.Threads)

	var wwg sync.WaitGroup
	var rewg sync.WaitGroup

	writer := bufio.NewWriter(os.Stdout)

	go wordDispenser(c.Wordlist, wordChan)

	for range c.Threads {
		wwg.Add(1)
		go c.worker(wordChan, resultChan, errorChan, &wwg, &rewg)
	}

	go resultWorker(resultChan, writer, &rewg)
	go c.errorWorker(errorChan, &rewg)

	wwg.Wait()
	rewg.Wait()
	close(resultChan)
	close(errorChan)
	sb := strings.Builder{}
	sb.WriteString("Total words fuzzed: ")
	sb.WriteString(strconv.Itoa(len(c.Wordlist)))
	sb.WriteString("\n")
	_, err := writer.WriteString(sb.String())
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (c Config) worker(wordChan <-chan string, resultChan chan<- string, errorChan chan<- string, wwg *sync.WaitGroup, rewg *sync.WaitGroup) {
	defer wwg.Done()

	for {
		word, ok := <-wordChan
		if !ok {
			return
		}

		result, err := c.Processor.ProcessWord(word)
		if err != nil {
			rewg.Add(1)
			errorChan <- err.Error()
			continue
		}

		if result == "" {
			continue
		}

		rewg.Add(1)
		resultChan <- result
	}

}

func wordDispenser(wordlist []string, wordChan chan<- string) {
	for _, word := range wordlist {
		wordChan <- word
	}

	close(wordChan)
}

func resultWorker(resultChan <-chan string, writer *bufio.Writer, rewg *sync.WaitGroup) {
	for {
		result, ok := <-resultChan
		if !ok {
			return
		}

		builder := strings.Builder{}

		builder.WriteString("Found: ")
		builder.WriteString(result)
		builder.WriteString("\n")

		writer.WriteString(builder.String())
		err := writer.Flush()
		if err != nil {
			panic(err)
		}

		rewg.Done()
	}
}

func (c Config) errorWorker(errorChan <-chan string, rewg *sync.WaitGroup) {
	for {
		error, ok := <-errorChan
		if !ok {
			return
		}
		c.Logger.Error(error)

		rewg.Done()
	}
}
