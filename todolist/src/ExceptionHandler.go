package src

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
    SERVER_ERROR    = 1000
    UNKNOWN_ERROR   = 1001
    PARAMETER_ERROR = 1003
)

type HandlerFunc func(c *gin.Context) error

type APIException struct {
    Code        int     `json:"-"`
    ErrorCode   int     `json:"error_code"`
    Msg         string  `json:"msg"`
    Request     string  `json:"request"`
}

func (e *APIException) Error() string {
    return  e.Msg
}

func newAPIException(code int, errorCode int,msg string) *APIException {
    return &APIException{
        Code:code,
        ErrorCode:errorCode,
        Msg:msg,
    }
}

func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
        var (
            err error
        )
        err = handler(c)
		if err != nil {
            var apiException *APIException
            if h,ok := err.(*APIException); ok {
                apiException = h
            }else if e, ok := err.(error); ok {
                if gin.Mode() == "debug" {
                    apiException = UnknownError(e.Error())
                }else{
                    apiException = UnknownError(e.Error())
                }
            }else{
                apiException = ServerError()
            }
            apiException.Request = c.Request.Method + " "+ c.Request.URL.String()
            c.JSON(apiException.Code,apiException)
            return
        }
    }
}

func ServerError() *APIException {
    return newAPIException(http.StatusInternalServerError,SERVER_ERROR,http.StatusText(http.StatusInternalServerError))
}

func UnknownError(message string) *APIException {
    return newAPIException(http.StatusForbidden,UNKNOWN_ERROR,message)
}

func ParameterError(message string) *APIException {
    return newAPIException(http.StatusBadRequest,PARAMETER_ERROR,message)
}