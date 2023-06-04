package errno

var (
	HTTPOK = &Errno{HTTPCode: 200, Code: "", Message: ""}

	InternalServerError = &Errno{HTTPCode: 500, Code: "InternalServerError", Message: "Internal server error."}

	SignTokenError = &Errno{HTTPCode: 401, Code: "AuthFailure.SignTokenError", Message: "Error occurred while signing the JSON web token."}

	TokenInvalidError = &Errno{HTTPCode: 401, Code: "AuthFailure.TokenInvalid", Message: "Token was invalid."}

	UnauthorizedError = &Errno{HTTPCode: 401, Code: "AuthFailure.Unauthorized", Message: "Unauthorized"}
)
