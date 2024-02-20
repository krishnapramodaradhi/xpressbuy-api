package customerror

const INTERNAL_SERVER_ERROR string = "An Internal Server Error has occured"

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Err        string `json:"error"`
	Message    string `json:"message"`
}

func New(code int, message string) error {
	return &CustomError{StatusCode: code, Err: getErrWithStatusCode(code), Message: message}
}

func (e *CustomError) Error() string {
	return e.Message
}

func getErrWithStatusCode(code int) string {
	switch code {
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	default:
		return "Internal Server Error"
	}
}
