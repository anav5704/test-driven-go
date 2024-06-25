package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")

func (d Dictionary) Search(query string) (string, error) {
	 definition, ok := d[query]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}


