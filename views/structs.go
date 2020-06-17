package views

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type PostRequest struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
