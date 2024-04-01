package register

import (
	db "simple_project/database"
	model "simple_project/models"
)


func Register(p model.Person) (ok bool) {
	if _, ok := db.Db[p.Email]; ok {
		return false
	}
	db.Db[p.Email] = p
	return true
}
