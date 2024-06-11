package commons

import (
	"github.com/saddam-satria/posq-be/domains/apis"
)



func GetResponse[T any](message string, statusCode int, data T) apis.Response[T]{
	response:= apis.Response[T]{
		Data: data,
		Message: message,
		StatusCode: statusCode,
	}


	return response
}	





