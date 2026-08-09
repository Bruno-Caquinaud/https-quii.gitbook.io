package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	missingWordError     = DictionaryErr("Missing Word in Dictionnary")
	duplicateWordError   = DictionaryErr("Insertion failed, word already exist")
	updateOperationError = DictionaryErr("Update operation error")
	deleteOperationError = DictionaryErr("Delete operation error")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, find := d[key]

	if !find {
		return "", missingWordError
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, find := d[key]

	if find {
		return duplicateWordError
	}
	d[key] = value

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case missingWordError:
		return updateOperationError
	case nil:
		d[key] = value

		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case missingWordError:
		return deleteOperationError
	case nil:
		delete(d, key)
		return nil
	default:
		return err
	}
}
