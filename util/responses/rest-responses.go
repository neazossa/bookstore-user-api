package responses

import (
	"time"
)

type Epoch int64

type Responses struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
	TimeServer Epoch       `json:"serverTime"`
}

func constructResponse(code string, message string, errors string, data interface{}) *Responses {
	var ep = time.Now().Unix()
	return &Responses{
		Status:     code,
		Message:    message,
		Error:      errors,
		Data:       data,
		TimeServer: Epoch(ep),
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
