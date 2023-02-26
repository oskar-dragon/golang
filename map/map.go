package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrWordNotFound = DictionaryErr("could not find a word")
	ErrWordAlreadyExist = DictionaryErr("word already exists")
	ErrWordCantBeUpdated = DictionaryErr("word cannot be updated because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]

	if !ok {
		return "",ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, definition string) error {
	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		d[key] = definition
	case nil:
		return ErrWordAlreadyExist
	default:
		return err;
	}

	return nil
}

func (d Dictionary) Update(key, newDefinition string) error {
	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		return ErrWordCantBeUpdated;
	case nil:
		d[key] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}