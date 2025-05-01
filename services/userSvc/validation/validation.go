package validation

import "net/mail"

type ValidationResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidEmailJson(email string) ValidationResponse {
	if IsValidEmail(email) {
		return ValidationResponse{
			Valid:   true,
			Message: "",
		}
	} else {
		return ValidationResponse{
			Valid:   false,
			Message: "Invalid Email Format",
		}
	}
}
