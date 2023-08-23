package main

import "errors"

type Dictionary map[string]string

var (
	ErrCantFinf   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Add(word, definition string) error {

	if _, err := d.Search(d[word]); err == nil {

		return ErrWordExists
	}
	d[word] = definition
	return nil
}
func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrCantFinf
	}

	return definition, nil
}
