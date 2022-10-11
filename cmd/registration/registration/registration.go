package registration

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"web-server/pkg/db/registrationdb"
)

func SignUp(login string, password string) (int64, error) {

	_, err := registrationdb.CheckCredentialsInDb(login)
	if err != nil {
		return 0, fmt.Errorf("CheckCredentialsInDb failed: %v", err)
	}

	_, err = VerifyLogin(login)
	if err != nil {
		return 0, fmt.Errorf("verifyLogin failed: %w", err)
	}

	if len(password) < 8 {
		return 0, fmt.Errorf("password must contain 8 characters")
	}

	err = CheckPassword(password)
	if err != nil {
		return 0, fmt.Errorf("error with password: %w", err)
	}

	hashPassword, passwordErr := hashPassword(password)
	if passwordErr != nil {
		return 0, fmt.Errorf("hashPassword failed: %w", passwordErr)
	}

	userId, err := registrationdb.CreateCredentials(context.Background(), login, hashPassword)
	if err != nil {
		return 0, fmt.Errorf("createCredentials failed: %w", err)
	}

	return userId, nil
}

func VerifyLogin(login string) (bool, error) {
	if len(login) < 8 {
		return false, fmt.Errorf("login must contain at least 8 characters")
	}

	number := `[0-9]{1}`
	a_z := `[a-z, A-Z]{1}`
	reg := `^[a-zA-Z0-9]+$`

	if b, err := regexp.MatchString(number, login); !b || err != nil {
		return false, fmt.Errorf("login need at least one number :%v", err)
	}

	if b, err := regexp.MatchString(a_z, login); !b || err != nil {
		return false, fmt.Errorf("login must contain at least one letter :%v", err)
	}

	if b, err := regexp.MatchString(reg, login); !b || err != nil {
		return false, fmt.Errorf("you have cyrillic")
	}

	return true, nil
}

func CheckPassword(ps string) error {
	if len(ps) < 8 {
		return fmt.Errorf("password len is < 8")
	}

	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`

	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("password need at least one number :%v", err)
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return fmt.Errorf("password must contain at least one letter in lower case :%v", err)
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("password must contain at least one letter in upper case :%v", err)
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("password need at least one special symbol :%v", err)
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
