package customerror

import "github.com/labstack/echo/v4"

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

func (e *CustomError) ErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	ce, ok := err.(*CustomError)
	if !ok {
		echoErr, _ := err.(*echo.HTTPError)
		c.JSON(echoErr.Code, New(echoErr.Code, err.Error()))
		return
	}
	if ce.StatusCode == 500 {
		ce.Message = INTERNAL_SERVER_ERROR
	}
	c.JSON(ce.StatusCode, ce)
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
