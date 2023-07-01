package maps

type Dictionary map[string]string

var (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]
	if !exists {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
