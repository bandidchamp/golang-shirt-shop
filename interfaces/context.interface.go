package interfaces

type HanbleError struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Request string      `json:"request"`
	Payload interface{} `json:"payload"`
}

type ResponsePayload struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

type RequestAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
