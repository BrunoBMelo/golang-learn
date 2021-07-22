package accounts

import (
	accounts "accounts/usecases/model"
)

func CreateNewAccount(username string, email string, password string) bool {

	if _, violations := accounts.BuildObject(username, email, password, true); violations == nil {
		return true
	}

	return false
}
