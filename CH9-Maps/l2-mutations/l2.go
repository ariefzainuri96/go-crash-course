package l2mutations

/*
Insert an Element

m[key] = elem
*/

/*
Get an Element

elem = m[key]
*/

/*
Delete an Element

delete(m, key)
*/

/*
Check If a Key Exists

elem, ok := m[key]
*/

import (
	"errors"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	elem, ok := users[name]

	// if maps with key name is not found
	if !ok {
		return false, errors.New("not found")
	}

	if !elem.scheduledForDeletion {
		return false, nil
	}

	// delete the maps with key name
	delete(users, name)

	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
