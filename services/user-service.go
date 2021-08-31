package services

import (
	"github.com/neazossa/bookstore-user-api/domain/users"
	"github.com/neazossa/bookstore-user-api/util/responses"
)

func CreateUser(u users.User) *responses.Responses {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.Save(); err != nil {
		return err
	}

	return responses.CreateSuccessResponse(&u, "")
}

func GetUser(u int64) *responses.Responses {
	result := users.User{Id: u}

	if err := result.Get(); err != nil {
		return err
	}

	return responses.CreateSuccessResponse(result, "")
}
