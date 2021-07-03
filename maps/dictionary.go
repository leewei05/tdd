package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("word already exists")
)

func (d Dictionary) Search(word string) (definition string, err error) {
	v, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return v, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}

	d[word] = definition
	return nil
}
