package helpers

import (
	"errors"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be blank")
	}
	if len(password) < 8 || len(password) > 30 {
		return errors.New("password length should be 8 to 30 characters")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("^[A-Za-z0-9$_@.#]+$"))) != nil {
		return errors.New("password should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("[0-9]"))) != nil {
		return errors.New("password should contain at least one number")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("[A-Za-z]"))) != nil {
		return errors.New("password should contain at least one alphabetic character")
	}
	return nil
}

func ValidateLogin(login string) error {
	if login == "" {
		return errors.New("login cannot be blank")
	}
	if len(login) < 5 || len(login) > 15 {
		return errors.New("login length should be 6 to 30 characters")
	}
	if validation.Validate(login, validation.Match(regexp.MustCompile("^[A-Za-z0-9$@_.#]+$"))) != nil {
		return errors.New("login should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	return nil
}

func ValidatePhone(phone string) error {
	re := regexp.MustCompile(`^[+][9][9][8]\d{9}$`)
	if !re.MatchString(phone) {
		return errors.New("invalid phone number")
	}

	return nil
}

func ValidateDates(startDate, endDate string) error {
	var layout = "2006-01-02 15:04:05"

	from, err := time.Parse(layout, startDate)
	if err != nil {
		return errors.New("start_date is invalid")
	}

	to, err := time.Parse(layout, endDate)
	if err != nil {
		return errors.New("end_time is invalid")
	}

	if !from.Before(to) {
		return errors.New("start_time can not be greater than end_time")
	}

	return nil
}
