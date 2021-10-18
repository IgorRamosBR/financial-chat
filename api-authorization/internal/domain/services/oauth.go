package services

import "api-authorization/internal/application/dto"

type OAuthService interface {
	Authorize(credentials dto.Auth)
}

type OAuthServiceImpl struct {
	userService UserService
}

func NewOAuthService(userService UserService) OAuthService {
	return OAuthServiceImpl{
		userService:  userService,
	}
}

func (s OAuthServiceImpl) Authorize(credentials dto.Auth) {
	//find user

	//check password

	//generate token

}