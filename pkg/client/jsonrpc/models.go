package jsonrpc

import "fmt"

type Request[T any] struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []T    `json:"params"`
	ID      int    `json:"id"`
}

type Response[T any] struct {
	ID      int            `json:"id"`
	Jsonrpc string         `json:"jsonrpc"`
	Result  T              `json:"result"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (e ErrorResponse) Error() error {
	return fmt.Errorf("code: %d, msg: %s, data: %#v", e.Code, e.Message, e.Data)
}

type CallParams struct {
	To   string `json:"to"`
	Data string `json:"data"`
}
