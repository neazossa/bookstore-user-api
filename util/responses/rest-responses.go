package responses

import (
	"github.com/neazossa/bookstore-user-api/util/date_utils"
)

type Responses struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
	TimeServer string      `json:"serverTime"`
}

func constructResponse(code string, message string, errors string, data interface{}) *Responses {
	return &Responses{
		Status:     code,
		Message:    message,
		Error:      errors,
		Data:       data,
		TimeServer: date_utils.GetNowString(),
	}
}

func CreateSuccessResponse(data interface{}, errors string) *Responses {
	return constructResponse("00", "Success", errors, data)
}

func CreateFailedResponse(code string, message string) *Responses {
	return constructResponse(code, message, "", "")
}

func CreateErrorResponses(message string) *Responses {
	return constructResponse("99", message, "", "")
}
