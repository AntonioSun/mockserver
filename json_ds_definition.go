package main

type RespMap []struct {
	HTTPRequest  HTTPRequest  `json:"httpRequest,omitempty"`
	HTTPResponse HTTPResponse `json:"httpResponse,omitempty"`
	// Times        Times        `json:"times,omitempty"`
	// TimeToLive   TimeToLive   `json:"timeToLive,omitempty"`
}

type HTTPRequest struct {
	Headers               Headers `json:"headers"`
	Method                string  `json:"method"`
	Path                  string  `json:"path"`
	QueryStringParameters AnyT    `json:"queryStringParameters"`
	Cookies               Cookies `json:"cookies"`
	Body                  Body    `json:"body"`
}

type Headers struct {
	ContentType []string `json:"Content-Type"`
}

type Body struct {
	Type        string `json:"type"`
	String      string `json:"string"`
	ContentType string `json:"contentType"`
}

type Cookies map[string]string
type AnyT map[string]interface{}

type HTTPResponse struct {
	StatusCode int     `json:"statusCode"`
	Body       string  `json:"body"`
	Cookies    Cookies `json:"cookies"`
}
