package data

import (
	"fmt"

	"github.com/AdieOlami/auth-service/src/errors"
	"github.com/AdieOlami/auth-service/src/model"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=? LIMIT 1 ALLOW FILTERING;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdate            = "UPDATE access_tokens SET expires=? WHERE accessToken=?;"
)

type TokenRepository interface {
	GetById(*gocql.Session, string) (*model.AccessToken, *errors.Error)
	Create(*gocql.Session, model.AccessToken) *errors.Error
	UpdateExpirationTime(*gocql.Session, model.AccessToken) *errors.Error
}

type tokenRepository struct {
}

func NewTokenRepository() TokenRepository {
	return &tokenRepository{}
}

func (r *tokenRepository) GetById(cassandra *gocql.Session, id string) (*model.AccessToken, *errors.Error) {

	var result model.AccessToken
	if err := cassandra.Query(queryGetAccessToken, id).Consistency(gocql.One).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found gor given id")
		}
		return nil, errors.NewInteralServerError(err.Error())
	}
	return &result, nil
}

func (r *tokenRepository) Create(cassandra *gocql.Session, at model.AccessToken) *errors.Error {
	fmt.Println("CREATE VALUES")
	fmt.Println(at.AccessToken)
	fmt.Println(at.UserId)
	fmt.Println(at.ClientId)
	fmt.Println(at.Expires)
	if err := cassandra.Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInteralServerError(err.Error())
	}
	return nil

}

func (r *tokenRepository) UpdateExpirationTime(cassandra *gocql.Session, at model.AccessToken) *errors.Error {

	if err := cassandra.Query(queryUpdate, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInteralServerError(err.Error())
	}
	return nil

}
