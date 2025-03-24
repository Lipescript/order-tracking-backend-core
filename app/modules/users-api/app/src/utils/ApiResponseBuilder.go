package utils

import (
	"users-api/app/src/constants"
	"users-api/app/src/domain"

)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constants.ResponseStatus, data T) domain.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) domain.ApiResponse[T] {
	return domain.ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
