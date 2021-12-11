package account

type AccountAuthenticationResponse struct {
	Token   string  `json:"token"`
	Profile Account `json:"profile"`
}

type EditAccountResponse struct {
	Ok    bool `json:"ok"`
	Message string `json:"message"`
}
