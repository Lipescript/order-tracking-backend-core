package constants

type ResponseStatus int
type Headers int
type General int

// logger
const (
	LoggerTimestampFormat                = "2006-01-02T15:04:05"
	LoggerLevelDebug                     = "DEBUG"
	LoggerLevelTrace                     = "TRACE"
	Success               ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
	MESSAGE_SUCCESSFUL_CONNECTION_TO_DATABASE = "Connected to database: %v successfully"
	MESSAGE_FAILED_CONNECTION_TO_DATABASE     = "Failed to connect to mongo database: %v. Disconnecting..."
)

// api & database
const (
	DefaultLocalPort           = ":8090"
	BaseAPI                    = "/backend/rgo/usersapi/v1/"
	DatabaseURI                = "mongodb://root:1234@localhost:27017"
	DatabaseName               = "rgo001_users_active_delivery"
	DatabaseTrackingCollection = "user_deliveries"
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}[r-1]
}
