package structs

type UserRegister struct {
	Name            string `json:"name" validate:"required"`
	MobileNumber    string `json:"mobile_number" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
