package pg_utils

import (
	"strings"

	"github.com/lib/pq"
	"github.com/neazossa/bookstore-user-api/util/responses"
)

const (
	noResultSet = "no rows in result set"
)

func ParseError(e error) *responses.Responses {
	sqlErr, ok := e.(*pq.Error)
	if !ok {
		if strings.Contains(e.Error(), noResultSet) {
			return responses.CreateFailedResponse("01", "no data found")
		}
		return responses.CreateErrorResponses("error parsing database responses")
	}

	switch sqlErr.Code {
	case "23505":
		return responses.CreateFailedResponse("01", "email already taken")
	}

	return responses.CreateErrorResponses("error processing request")
}
