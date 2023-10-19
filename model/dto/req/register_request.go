package req

type RegisterRequest struct {
	Identifier      registerIdentifier
	UserName        string `json:"userName" validate:"required,min=3,max=30"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,eqfield=Password"`
}

type registerIdentifier struct {
	Email       string `json:"email" validate:"email"`
	PhoneNumber string `json:"phoneNumber" validate:"min=10,max=15"`
}
