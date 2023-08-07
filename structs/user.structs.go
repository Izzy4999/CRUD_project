package structs

type Register struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,min=8,max=12"`
	Password    string `json:"password" validate:"required,min=6,max=8"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Verify struct {
	Code  string `json:"code" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
