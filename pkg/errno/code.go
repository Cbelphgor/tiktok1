package errno

import code "github.com/a76yyyy/ErrnoCode"

const (
	// ErrVideoNotFound - 400: Video not found.
	codeErrVideoNotFound int = iota + 110201
)

var (
	HttpSuccess              = NewHttpErr(code.ErrSuccess, 200, "OK")
	ErrHttpUnknown           = NewHttpErr(code.ErrUnknown, 500, "Internal server error")
	ErrHttpBind              = NewHttpErr(code.ErrBind, 400, "Error occurred while binding the request body to the struct")
	ErrHttpValidation        = NewHttpErr(code.ErrValidation, 400, "Validation failed")
	ErrHttpTokenInvalid      = NewHttpErr(code.ErrTokenInvalid, 401, "Token invalid")
	ErrHttpDatabase          = NewHttpErr(code.ErrDatabase, 500, "Database error")
	ErrHttpEncrypt           = NewHttpErr(code.ErrEncrypt, 401, "Error occurred while encrypting the user password")
	ErrHttpSignatureInvalid  = NewHttpErr(code.ErrSignatureInvalid, 401, "Signature is invalid")
	ErrHttpExpired           = NewHttpErr(code.ErrExpired, 401, "Token expired")
	ErrHttpInvalidAuthHeader = NewHttpErr(code.ErrInvalidAuthHeader, 401, "Invalid authorization header")
	ErrHttpMissingHeader     = NewHttpErr(code.ErrMissingHeader, 401, "The `Authorization` header was empty")
	ErrHttporExpired         = NewHttpErr(code.ErrorExpired, 401, "Token expired")
	ErrHttpPasswordIncorrect = NewHttpErr(code.ErrPasswordIncorrect, 401, "Password was incorrect")
	ErrHttpPermissionDenied  = NewHttpErr(code.ErrPermissionDenied, 403, "Permission denied")
	ErrHttpEncodingFailed    = NewHttpErr(code.ErrEncodingFailed, 500, "Encoding failed due to an error with the data")
	ErrHttpDecodingFailed    = NewHttpErr(code.ErrDecodingFailed, 500, "Decoding failed due to an error with the data")
	ErrHttpInvalidJSON       = NewHttpErr(code.ErrInvalidJSON, 500, "Data is not valid JSON")
	ErrHttpEncodingJSON      = NewHttpErr(code.ErrEncodingJSON, 500, "JSON data could not be encoded")
	ErrHttpDecodingJSON      = NewHttpErr(code.ErrDecodingJSON, 500, "JSON data could not be decoded")
	ErrHttpInvalidYaml       = NewHttpErr(code.ErrInvalidYaml, 500, "Data is not valid Yaml")
	ErrHttpEncodingYaml      = NewHttpErr(code.ErrEncodingYaml, 500, "Yaml data could not be encoded")
	ErrHttpDecodingYaml      = NewHttpErr(code.ErrDecodingYaml, 500, "Yaml data could not be decoded")
	ErrHttpUserNotFound      = NewHttpErr(code.ErrUserNotFound, 404, "User not found")
	ErrHttpUserAlreadyExist  = NewHttpErr(code.ErrUserAlreadyExist, 400, "User already exist")
	ErrHttpReachMaxCount     = NewHttpErr(code.ErrReachMaxCount, 400, "Secret reach the max count")
	ErrHttpSecretNotFound    = NewHttpErr(code.ErrSecretNotFound, 404, "Secret not found")
)

var (
	Success              = NewErrNo(code.ErrSuccess, "OK")
	ErrUnknown           = NewErrNo(code.ErrUnknown, "Internal server error")
	ErrBind              = NewErrNo(code.ErrBind, "Error occurred while binding the request body to the struct")
	ErrValidation        = NewErrNo(code.ErrValidation, "Validation failed")
	ErrTokenInvalid      = NewErrNo(code.ErrTokenInvalid, "Token invalid")
	ErrDatabase          = NewErrNo(code.ErrDatabase, "Database error")
	ErrEncrypt           = NewErrNo(code.ErrEncrypt, "Error occurred while encrypting the user password")
	ErrSignatureInvalid  = NewErrNo(code.ErrSignatureInvalid, "Signature is invalid")
	ErrExpired           = NewErrNo(code.ErrExpired, "Token expired")
	ErrInvalidAuthHeader = NewErrNo(code.ErrInvalidAuthHeader, "Invalid authorization header")
	ErrMissingHeader     = NewErrNo(code.ErrMissingHeader, "The `Authorization` header was empty")
	ErrorExpired         = NewErrNo(code.ErrorExpired, "Token expired")
	ErrPasswordIncorrect = NewErrNo(code.ErrPasswordIncorrect, "Password was incorrect")
	ErrPermissionDenied  = NewErrNo(code.ErrPermissionDenied, "Permission denied")
	ErrEncodingFailed    = NewErrNo(code.ErrEncodingFailed, "Encoding failed due to an error with the data")
	ErrDecodingFailed    = NewErrNo(code.ErrDecodingFailed, "Decoding failed due to an error with the data")
	ErrInvalidJSON       = NewErrNo(code.ErrInvalidJSON, "Data is not valid JSON")
	ErrEncodingJSON      = NewErrNo(code.ErrEncodingJSON, "JSON data could not be encoded")
	ErrDecodingJSON      = NewErrNo(code.ErrDecodingJSON, "JSON data could not be decoded")
	ErrInvalidYaml       = NewErrNo(code.ErrInvalidYaml, "Data is not valid Yaml")
	ErrEncodingYaml      = NewErrNo(code.ErrEncodingYaml, "Yaml data could not be encoded")
	ErrDecodingYaml      = NewErrNo(code.ErrDecodingYaml, "Yaml data could not be decoded")

	ErrUserNotFound     = NewErrNo(code.ErrUserNotFound, "User not found")
	ErrUserAlreadyExist = NewErrNo(code.ErrUserAlreadyExist, "User already exist")
	ErrReachMaxCount    = NewErrNo(code.ErrReachMaxCount, "Secret reach the max count")
	ErrSecretNotFound   = NewErrNo(code.ErrSecretNotFound, "Secret not found")
	ErrVideoNotFound    = NewErrNo(codeErrVideoNotFound, "Video not found")
)
