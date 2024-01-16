package requests

/*
Defining all auth related api request structs
In the validate fields '@' symbol has been used as a field params separator
*/
type UserRegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	MobileNumber    string `json:"mobile_number" validate:"required,unique=golist_users@mobile_number,len=12"`
	Email           string `json:"email" validate:"required,email,unique=golist_users@email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UserLoginRequest struct {
	Username string `josn:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
