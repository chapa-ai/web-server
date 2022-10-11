package models

type Error struct {
	Code int    `json:"code,omitempty"`
	Text string `json:"text,omitempty"`
}

type Response struct {
	Response interface{} `json:"response,omitempty"`
	Login    string      `json:"login,omitempty"`
}

type Model struct {
	Error    *Error      `json:"error,omitempty"`
	Response *Response   `json:"response,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
