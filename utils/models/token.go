package models

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUuid   string `json:"-"`
	RefreshUuid  string `json:"-"`
	ATExpiresAt  int64  `json:"-"`
	RTExpiresAt  int64  `json:"-"`
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}
