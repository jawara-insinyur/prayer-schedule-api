package handler

import "encoding/json"

type clientError interface {
	Error() string
	ResponseBody() ([]byte, error)
	ResponseHeader() (int, map[string]string)
}

type httpError struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Cause   error  `json:"-"`
}

func (e *httpError) Error() string {
	return e.Cause.Error()
}

func (e *httpError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (e *httpError) ResponseHeader() (int, map[string]string) {
	return e.Status, map[string]string{"Content-Type": "application/json"}
}

func NewHttpError(err error, status int, message string) error {
	return &httpError{Cause: err, Status: status, Message: message}
}
