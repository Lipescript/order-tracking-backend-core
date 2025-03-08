package constants

const configConstants {
		loggerTimestampFormat = "2006-01-02T15:04:05"
		loggerLevelDebug = "DEBUG"
		loggerLevelTrace = "TRACE"
}

type ResponseStatus int
type Headers int
type General int

// Constant Api
const (
   Success ResponseStatus = iota + 1
   DataNotFound
   UnknownError
   InvalidRequest
   Unauthorized
)

func (r ResponseStatus) GetResponseStatus() string {
   return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
   return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}[r-1]
}