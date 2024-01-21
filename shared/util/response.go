package util

import "system-point/delivery/http/dto/response"

func ConstructResponseError(statusCode int, errorMsg string) (response.ResponseErrorDTO, int) {
	resp := response.ResponseErrorDTO{
		StatusCode: statusCode,
		Error:      errorMsg,
	}

	return resp, statusCode
}
