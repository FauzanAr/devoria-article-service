package account

// AccountRegistrationRequest is a model for account registration.
type AccountRegistrationRequest struct {
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

// AccountAuthenticationRequest is a model of account authentication.
type AccountAuthenticationRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AccountUpdateRequest struct {
	ID             int64      `json:"id"`
	Password       string    `json:"password"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
}
