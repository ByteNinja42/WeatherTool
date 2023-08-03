package service

type ErrorResponseAPI struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	StatusCode int `json:"statusCode"`
}
