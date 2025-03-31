package utils

import (
	"users-api/app/src/constants"
	"users-api/app/src/domain"

)

func Null() interface{} {
	return nil
}


func BuildResponse_[T any](responseStatus constants.ResponseStatus, message string, data T) domain.ApiResponse[T] {
	return domain.ApiResponse[T]{
		ResponseKey:     responseStatus.GetResponseStatus(),
		ResponseMessage: message,
		Data:            data,
	}
}
