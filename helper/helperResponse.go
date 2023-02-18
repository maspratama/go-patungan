package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	ResponseCode    int    `json:"responseCode"`
	ResponseStatus  string `json:"responseStatus"`
	ResponseMessage string `json:"responseMessage"`
}

func APIResponse(responseMessage string, responseCode int, responseStatus string, data interface{}) Response {
	meta := Meta{
		ResponseCode:    responseCode,
		ResponseStatus:  responseStatus,
		ResponseMessage: responseMessage,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
