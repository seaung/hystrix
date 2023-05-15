package errno

var (
	HTTPOK = &Errno{HTTPCode: 200, Code: "", Message: ""}

	InternalServerError = &Errno{HTTPCode: 500, Code: "InternalServerError", Message: "Internal server error."}
)
