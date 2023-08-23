package main

import "errors"

type Dictionary map[string]string

var ErrCantFinf = "could not find the word you were looking for"

func (d Dictionary) Add(word, definition string) {

	d[word] = definition

}
func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", errors.New(ErrCantFinf)
	}

	return definition, nil
}
