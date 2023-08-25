package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrCantFinf         = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Update(word, def string) error {

	if _, err := d.Search(word); err != nil {
		return ErrWordDoesNotExist
	}
	d[word] = def
	return nil
}

func (d Dictionary) Add(word, definition string) error {

	if _, err := d.Search(word); err == nil {

		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Add2(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrCantFinf
	}

	return definition, nil
}
