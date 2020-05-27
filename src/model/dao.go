package model

// import (
// 	"github.com/gocql/gocql"

// 	"github.com/AdieOlami/auth-service/src/errors"
// )

// const (
// 	queryGetAccessToken    = "SELECT accessToken, userId, clientId, expires FROM accessTokens WHERE accessToken=?;"
// 	queryCreateAccessToken = "INSERT INTO accessTokens(accessToken, userId, clientId, expires) VALUES (?, ?, ?, ?);"
// 	queryUpdate            = "UPDATE accessTokens SET expires=? WHERE accessToken=?;"
// )

// func (at *AccessToken) GetById(cassandra *gocql.Session, id string) (*AccessToken, *errors.Error) {

// 	var result AccessToken
// 	if err := cassandra.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
// 		if err == gocql.ErrNotFound {
// 			return nil, errors.NewNotFoundError("no access token found gor given id")
// 		}
// 		return nil, errors.NewInteralServerError(err.Error())
// 	}
// 	return &result, nil
// }

// func (at *AccessToken) Create(cassandra *gocql.Session) *errors.Error {

// 	if err := cassandra.Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
// 		return errors.NewInteralServerError(err.Error())
// 	}
// 	return nil

// }

// func (at *AccessToken) UpdateExpirationTime(cassandra *gocql.Session) *errors.Error {

// 	if err := cassandra.Query(queryUpdate, at.Expires, at.AccessToken).Exec(); err != nil {
// 		return errors.NewInteralServerError(err.Error())
// 	}
// 	return nil

// }
