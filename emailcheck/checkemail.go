package emailcheck

import (
	db "simple_project/database"
	model "simple_project/models"
)

func CHeckEmail(email string) (model.Person, bool) {
	a, b := db.Db[email]
	return a, b
}
