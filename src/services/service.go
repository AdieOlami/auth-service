package services

import (
	"fmt"
	"strings"

	"github.com/AdieOlami/auth-service/src/data"
	"github.com/AdieOlami/auth-service/src/errors"
	"github.com/AdieOlami/auth-service/src/model"
	"github.com/gocql/gocql"
)

type Service interface {
	GetById(string) (*model.AccessToken, *errors.Error)
	Create(model.AccessTokenRequest) (*model.AccessToken, *errors.Error)
	UpdateExpirationTime(model.AccessToken) *errors.Error
}

type service struct {
	cassandra    *gocql.Session
	tokenService data.TokenRepository
	userService  data.RestUsersRepository
}

func NewService(cassandra *gocql.Session, tokenService data.TokenRepository, userService data.RestUsersRepository) Service {
	return &service{
		cassandra:    cassandra,
		tokenService: tokenService,
		userService:  userService,
	}
}

func (s *service) GetById(accessTokenId string) (*model.AccessToken, *errors.Error) {
	// at := &model.AccessToken{AccessToken: accessTokenId}
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.tokenService.GetById(s.cassandra, accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(req model.AccessTokenRequest) (*model.AccessToken, *errors.Error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.userService.LoginUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	fmt.Println("THE USER %s", user)

	// Generate a new access token:
	at := model.GetNewAccessToken(user.ID.String())
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.tokenService.Create(s.cassandra, at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at model.AccessToken) *errors.Error {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.tokenService.UpdateExpirationTime(s.cassandra, at)
	// at.UpdateExpirationTime(s.cassandra)
}
