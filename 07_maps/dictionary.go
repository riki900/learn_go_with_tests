package main

import "errors"

var ErrNotFound = errors.New("cannot find word in dictionary")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
