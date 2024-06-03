package main

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrWordNotFound            = DictionaryErr("could not find the word you were looking for")
	ErrWordExists              = DictionaryErr("cannot add word, it already exists")
	ErrUpdatedWordDoesNotExist = DictionaryErr("cannot update word, it doesn't exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, val string) error {
	_, ok := d[key]

	if ok {
		return ErrWordExists
	}

	d[key] = val

	return nil
}

func (d Dictionary) Update(key, val string) error {
	_, ok := d[key]

	if !ok {
		return ErrUpdatedWordDoesNotExist
	}

	d[key] = val

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
