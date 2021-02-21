package helpers


type ResponseSuccessStruct struct {
	Status string `json:"status"`
	Code   int	  `json:"code"`
	Data   interface{} `json:"data"`
}

type ResponseFailedStruct struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Message interface{} `json:"message"`
}

func ResponseSuccess(status string, code int, data interface{}) ResponseSuccessStruct {
	jsonResponse := ResponseSuccessStruct{
		Status: status,
		Code: code,
		Data: data,
	}

	return jsonResponse
}

func ResponseFailed(status string, code int, message interface{}) ResponseFailedStruct {
	jsonResponse := ResponseFailedStruct{
		Status: status,
		Code: code,
		Message: message,
	}

	return jsonResponse
}