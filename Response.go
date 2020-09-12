package main

type Response struct {
	Datas ResponseData `json:"datas"`
}

type ResponseData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	ID    string `json:"id"`
	Image string `json:"image"`
}
type Token struct {
	AccessToken string `json:"access_token"`
}

type ErrorResponse struct {
	Error        string      `json:"error"`
	ErrorCode    interface{} `json:"error_code"`
	ErrorMessage interface{} `json:"error_messages"`
}
