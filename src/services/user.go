package services

// import (
// 	"encoding/json"

// 	"github.com/AdieOlami/auth-service/src/errors"
// 	"github.com/AdieOlami/auth-service/src/model"
// )

// type UserService interface {
// 	LoginUser(string, string) (*model.User, *errors.Error)
// }

// type userService struct{}

// func NewUserService() UserService {
// 	return &userService{}
// }

// func (r *userService) LoginUser(email string, password string) (*model.User, *errors.Error) {
// 	request := model.UserLoginRequest{
// 		Email:    email,
// 		Password: password,
// 	}

// 	response := usersRestClient.Post("/users/login", request)

// 	if response == nil || response.Response == nil {
// 		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))
// 	}

// 	if response.StatusCode > 299 {
// 		apiErr, err := errors.NewRestErrorFromBytes(response.Bytes())
// 		if err != nil {
// 			return nil, errors.NewInternalServerError("invalid error interface when trying to login user", err)
// 		}
// 		return nil, apiErr
// 	}

// 	var user model.User
// 	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
// 		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response", errors.New("json parsing error"))
// 	}
// 	return &user, nil
// }
