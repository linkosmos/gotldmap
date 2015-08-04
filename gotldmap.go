package gotldmap

import "errors"

// ErrTLDNotFound -
var ErrTLDNotFound = errors.New("TLD not found")

// TldExist - checks whether given TLD is in the map
func TldExist(tld string) bool {
	_, exist := Map[tld]
	return exist
}

// FindByTld - returns mark of tld
func FindByTld(tld string) (int, error) {
	if value, ok := Map[tld]; ok {
		return value, nil
	}
	return 0, ErrTLDNotFound
}

// FindByMark - returns tld of corresponding mark
func FindByMark(mark int) (string, error) {
	for k, v := range Map {
		if v == mark {
			return k, nil
		}
	}
	return "", ErrTLDNotFound
}
