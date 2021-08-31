package users

import "github.com/neazossa/bookstore-user-api/util/responses"

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (user *User) Validate() *responses.Responses {
	if user.Email == "" {
		return responses.CreateFailedResponse("01", "invalid email address")
	}
	if user.FirstName == "" {
		return responses.CreateFailedResponse("01", "invalid firstname")
	}
	if user.LastName == "" {
		return responses.CreateFailedResponse("01", "invalid lastname")
	}
	if user.Username == "" {
		return responses.CreateFailedResponse("01", "invalid username")
	}
	if user.Password == "" {
		return responses.CreateFailedResponse("01", "invalid password")
	}

	return nil
}
