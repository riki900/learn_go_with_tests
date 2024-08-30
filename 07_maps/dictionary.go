package main

const (
	ErrNotFound          = DictionaryErr("cannot find word in dictionary")
	ErrWordExists        = DictionaryErr("cannod add existing word")
	ErrWordDoesNotExists = DictionaryErr("operation failed: word not in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {

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

func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrWordDoesNotExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {

	delete(d, word)

}
