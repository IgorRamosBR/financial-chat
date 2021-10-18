package dto

import "errors"

const PASSWORD_GRANT_TYPE = "password"

type Auth struct {
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (a Auth) IsValid() error {
	if a.GrantType != PASSWORD_GRANT_TYPE {
		return errors.New("Invalid grant type")
	}
	return nil
}