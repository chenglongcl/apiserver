package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation     = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase       = &Errno{Code: 20002, Message: "Database error."}
	ErrToken          = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrRecordNotFound = &Errno{Code: 20006, Message: "record not found"}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrUserExist         = &Errno{Code: 20105, Message: "User already exist"}
	ErrNotUserExist      = &Errno{Code: 20106, Message: "User does not exist"}

	//upload errors
	ErrUploadFile               = &Errno{Code: 20201, Message: "Error uploadFile"}
	ErrUploadMime               = &Errno{Code: 20202, Message: "Error uploadMime"}
	ErrUploadFail               = &Errno{Code: 20203, Message: "Upload fail"}
	ErrOssGenerateSignatureFail = &Errno{Code: 20204, Message: "AliYunOss Signature fail"}
	ErrAliYunBucket             = &Errno{Code: 20205, Message: "阿里云OSS Bucket读取失败"}
	ErrAliYunOssUploadFail      = &Errno{Code: 20205, Message: "阿里云OSS上传失败"}
)
