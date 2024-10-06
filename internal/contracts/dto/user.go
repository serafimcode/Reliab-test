package dto

import _ "github.com/go-playground/validator/v10"

/*TODO check validation*/
type UserDTO struct {
	Type  string `json:"string" validate:"oneof=USER APPLICANT"`
	Email string `json:"email"`
	FIO   string `json:"fio"`
}
