package fuzzlib

type Processor interface {
	ProcessWord(word string) (string, error)
}
