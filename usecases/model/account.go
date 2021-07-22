package accounts

import (
	"net/mail"
	"regexp"
)

type account struct {
	Id       int
	Username string
	Email    string
	Password string
	Actived  bool
}

const (
	msgEmailIsNil                  = "E-mail cannot be null or empty"
	msgEmailInvalid                = "E-mail is invalid"
	msgUsernameInvalid             = "Username cannot be null or empty"
	msgPasswordInvalid             = "Password cannot be null or empty"
	msgPasswordLenghtInvalid       = "Password must be 8 position"
	msgPasswordWithoutSpecialChars = "Password must contain at least 1 special character"
	msgPasswordWithoutUpperCase    = "Password must be capital letter"
	msgPasswordWithoutLowerCase    = "Password must be lower case"
	msgPasswordWithoutNumber       = "Password must be numbers"
)

func BuildObject(username string, email string, password string, actived bool) (objBuilder account, violations []string) {

	if emailIsNil := valueIsNil(email); !emailIsNil {
		violations = append(violations, msgEmailIsNil)
	}

	if emailIsValid := emailIsValid(email); !emailIsValid {
		violations = append(violations, msgEmailInvalid)
	}

	if pwdValidationResult := checkIfPasswordIsLegible(password); len(pwdValidationResult) > 0 {
		violations = append(violations, pwdValidationResult...)
	}

	if usernameIsValid := valueIsNil(username); !usernameIsValid {
		violations = append(violations, msgUsernameInvalid)
	}

	if passwordIsValid := valueIsNil(password); !passwordIsValid {
		violations = append(violations, msgPasswordInvalid)
	}

	if hasViolations(violations) {
		return account{}, violations
	}

	return account{0, username, email, password, actived}, nil
}

func checkIfPasswordIsLegible(value string) (validation []string) {

	if hasSpecialCharacterer, _ := regexp.MatchString(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]`, value); !hasSpecialCharacterer {
		validation = append(validation, msgPasswordWithoutSpecialChars)
	}

	if hasCapitalLetter, _ := regexp.MatchString(`[A-Z]`, value); !hasCapitalLetter {
		validation = append(validation, msgPasswordWithoutUpperCase)
	}

	if hasLowerCase, _ := regexp.MatchString(`[a-z]`, value); !hasLowerCase {
		validation = append(validation, msgPasswordWithoutLowerCase)
	}

	if hasNumber, _ := regexp.MatchString(`[0-9]`, value); !hasNumber {
		validation = append(validation, msgPasswordWithoutNumber)
	}

	if lengthIsValid := passwordHasLenthValid(value); !lengthIsValid {
		validation = append(validation, msgPasswordLenghtInvalid)
	}

	return validation
}

func valueIsNil(value string) bool {
	return len(value) > 0
}

func hasViolations(value []string) bool {
	return len(value) > 0
}

func emailIsValid(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}

func passwordHasLenthValid(value string) bool {
	return len(value) == 8
}
