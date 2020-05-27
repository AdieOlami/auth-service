package data

import (
	"encoding/json"
	"fmt"

	"github.com/AdieOlami/auth-service/src/errors"
	"github.com/AdieOlami/auth-service/src/model"
	"github.com/go-resty/resty/v2"
)

var (
	restClient = &resty.Client{
		HostURL: "http://localhost:9090",
	}
	// usersRestClient = rest.RequestBuilder{
	// 	BaseURL: "http://localhost:9090",
	// 	Timeout: 100 * time.Millisecond,
	// }
)

type RestUsersRepository interface {
	LoginUser(string, string) (*model.User, *errors.Error)
}

type usersRepository struct{}

func NewUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*model.User, *errors.Error) {
	request := model.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	var user model.User

	var respData model.ResponseData
	client := resty.New()
	client.SetHostURL("http://localhost:9090")

	resp, err := client.R().
		SetBody(request).
		SetResult(&respData). // or SetResult(AuthSuccess{}).
		Post("/api/v1/users/login")

	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("invalid restclient response when trying to login user %s", err.Error()))
	}

	jsonErr := json.Unmarshal(resp.Body(), &user)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return nil, errors.NewInternalServerError(fmt.Sprintf("error decoding json %s", jsonErr.Error()))
	}

	if resp.StatusCode() != 200 {
		return nil, errors.NewInternalServerError(fmt.Sprintf("invalid restclient invalid status code %s", err.Error()))
	}
	return &respData.Data, nil
}
