package main

import "errors"

var ErrWordNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(str string) (string, error) {
	definition, ok := d[str]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, val string) {
	d[key] = val
}
