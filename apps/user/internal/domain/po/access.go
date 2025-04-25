package po

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	VerCode  string `json:"verCode"`
	VerToken string `json:"verToken"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
}
