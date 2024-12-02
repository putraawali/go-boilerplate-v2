package response

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

type Response struct{}

func NewResponse() *Response {
	return &Response{}
}

type response struct {
	Error   any    `json:"error"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (r *Response) Send(statusCode int, data any, err any) (int, *response) {
	resp := &response{}
	if err != nil {
		resp.Message = "error"

		errData, ok := err.(*ErrorResponse)
		if ok {
			resp.Error = errorResponse{
				StatusCode: errData.GetStatusCode(),
				Title:      errData.GetTitle(),
				Message:    errData.GetMessage(),
				Detail:     errData.GetDetail(),
				Source:     errData.GetSource(),
			}

			statusCode = errData.statusCode
			errData.Log()
		} else {
			// Generate stack trace
			source := ""

			pc, file, line, ok := runtime.Caller(1)
			funcName := ""

			if details := runtime.FuncForPC(pc); details != nil {
				titles := strings.Split(details.Name(), ".")
				funcName = titles[len(titles)-1]
			}

			if ok {
				source = fmt.Sprintf("Called from %s, line #%d, func: %v", file, line, funcName)
			}

			resp.Error = errorResponse{
				Detail:     fmt.Sprint(err),
				Title:      http.StatusText(http.StatusInternalServerError),
				Message:    http.StatusText(http.StatusInternalServerError),
				Source:     source,
				StatusCode: http.StatusInternalServerError,
			}

			statusCode = http.StatusInternalServerError
		}
	} else {
		resp.Data = data
		resp.Message = "success"
	}

	return statusCode, resp
}
