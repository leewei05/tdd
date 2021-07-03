package dictionary

type Dictionary map[string]string

type DictionaryErr string

var (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("word already exists")
	ErrWordDoesNotExists = DictionaryErr("word does not exist")
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

func (d Dictionary) Update(word, newDefinition string) error {
	_, ok := d[word]
	if !ok {
		return ErrWordDoesNotExists
	}

	d[word] = newDefinition
	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
