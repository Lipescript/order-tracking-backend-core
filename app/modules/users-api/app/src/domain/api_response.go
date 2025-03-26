package domain

type ApiResponse[T any] struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
	Data            T      `json:"data"`
}

// POST /users successfully created users.
func UserCreatedResponse[T any](data T) ApiResponse[T] {
	return ApiResponse[T]{
		ResponseKey:     "USER_CREATED",
		ResponseMessage: "User successfully created in the database",
		Data:            data,
	}
}
