package services

import "errors"

var (
	noResponseToParse = errors.New("no_response_to_parse")
	missingRecaptchaToken = errors.New("token_cannot_be_nil_for_recap_type")
)