package response

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

type ErrorResponse struct {
	statusCode int
	detail     string
	title      string
	message    string
	source     string
	ctx        context.Context
	err        error
	stack      []string
	fn         string
	line       int
	path       string
}

type errorResponse struct {
	StatusCode int    `json:"status_code"`
	Title      string `json:"title"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
	Source     string `json:"source"`
}

// implement interface error
func (e *ErrorResponse) Error() string {
	return e.err.Error()
}

func (r *Response) NewError() *ErrorResponse {
	e := &ErrorResponse{}

	e.source = ""

	// Get stack trace when NewError function called
	pc, file, line, ok := runtime.Caller(1)
	funcName := ""

	if details := runtime.FuncForPC(pc); details != nil {
		titles := strings.Split(details.Name(), ".")
		funcName = titles[len(titles)-1]
	}

	if ok {
		e.source = fmt.Sprintf("Called from %s, line #%d, func: %v", file, line, funcName)
	}

	e.stack = trace(2)

	e.fn = funcName
	e.line = line
	e.path = file

	return e
}

func (e *ErrorResponse) SetStatusCode(code int) *ErrorResponse {
	e.statusCode = code
	e.title = http.StatusText(code)
	return e
}

func (e *ErrorResponse) SetDetail(detail string) *ErrorResponse {
	e.detail = detail
	return e
}

func (e *ErrorResponse) SetMessage(message error) *ErrorResponse {
	e.message = message.Error()
	e.err = message

	return e
}

func (e *ErrorResponse) SetContext(ctx context.Context) *ErrorResponse {
	e.ctx = ctx
	return e
}

func (e *ErrorResponse) GetStatusCode() int {
	return e.statusCode
}

func (e *ErrorResponse) GetTitle() string {
	return e.title
}

func (e *ErrorResponse) GetMessage() string {
	return e.message
}

func (e *ErrorResponse) GetDetail() string {
	return e.detail
}

func (e *ErrorResponse) GetSource() string {
	return e.source
}
