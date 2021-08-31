package users

import "github.com/neazossa/bookstore-user-api/util/responses"

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *responses.Responses {
	result := userDB[user.Id]
	if result == nil {
		return responses.CreateFailedResponse("02", "user not found")
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Username = result.Username
	user.Password = result.Password
	user.Email = result.Email
	return nil
}

func (user *User) Save() *responses.Responses {
	currentUser := userDB[user.Id]
	if currentUser != nil {
		if currentUser.Email != user.Email {
			return responses.CreateFailedResponse("01", "email already registered")
		}

		return responses.CreateFailedResponse("01", "user already exists")
	}
	userDB[user.Id] = user
	return nil
}
