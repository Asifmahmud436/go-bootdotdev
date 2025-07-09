package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	for k, v := range users {
		if v.name == name {
			if v.scheduledForDeletion == true{
				delete(users,k)
				return true, nil
			}
			return false, nil
		}
	}
	return false, errors.New("not found")
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
