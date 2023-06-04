package errno

var (
	UserAlreadyExistError = &Errno{HTTPCode: 400, Code: "OperationFailed.UserAlreadyExist", Message: "User already exist"}

	UserNotFoundError = &Errno{HTTPCode: 400, Code: "ResourceNotFound.UserNotFound", Message: "User was not found"}

	PasswordIncorrectError = &Errno{HTTPCode: 401, Code: "InvalidParameter.PasswordIncorrect", Message: "Password was incorrect"}
)
