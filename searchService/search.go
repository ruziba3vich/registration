package searchservice

import (
	db "simple_project/database"
	m "simple_project/models"
)

func Search(name string, surname string) (bool, m.Person) {
	dummy := m.NewPerson("", "", 0, false, "", "", "", "")
	for _, person := range db.Db {
		if person.Name == name && person.Surname == surname {
			return true, person
		}
	}
	return false, dummy
}
