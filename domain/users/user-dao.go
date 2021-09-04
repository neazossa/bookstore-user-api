package users

import (
	"github.com/neazossa/bookstore-user-api/datasources/user_db"
	"github.com/neazossa/bookstore-user-api/util/hash_utils"
	"github.com/neazossa/bookstore-user-api/util/pg_utils"
	"github.com/neazossa/bookstore-user-api/util/responses"
)

const (
	indexEmail      = "m_user_index_2"
	queryInsertUser = "INSERT INTO m_user(firstname, lastname, username, email, userpass) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	queryGetUser    = "SELECT id, firstname, lastname, username, email, userpass from m_user where id = $1"
)

func (user *User) Get() *responses.Responses {

	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		return responses.CreateErrorResponses(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password); err != nil {
		return pg_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *responses.Responses {
	var id int64

	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return responses.CreateErrorResponses(err.Error())
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.FirstName, user.LastName, user.Username, user.Email, hash_utils.HashPassword(user.Password)).Scan(&id)
	if err != nil {
		return pg_utils.ParseError(err)
	}

	user.Id = int64(id)
	return nil
}
