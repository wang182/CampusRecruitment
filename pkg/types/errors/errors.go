package errors

var (
	// error code 组成
	// - 第一位：一级分类
	// - 第二三位：二级分类
	// - 第四五六位: 分类下的错误码

	// 通用错误
	ErrUnknown        = newErrorCode("unknown error", 100100)
	ErrDBError        = newErrorCode("database error", 100200)
	ErrNotImplemented = newErrorCode("not implemented", 100310)
	ErrBadRequest     = newErrorCode("bad request", 100400)
	ErrNotAuth        = newErrorCode("authentication required", 100430)
	ErrAuthFailed     = newErrorCode("authentication failed", 100431)
	ErrStorageError   = newErrorCode("storage error", 100500)
	ErrEncrypt        = newErrorCode("encrypt error", 100600)

	// not found error
	// 此类型 error 必须为 104xxx，ErrNotFound 必须为 104000，以支持统一的 IsNotFoundErr() 判断
	ErrNotFound     = newErrorCode("not found", 104000)
	ErrUserNotFound = newErrorCode("user not found", 104060)

	// 权限错误
	ErrPermDeny = newErrorCode("permission denied", 110100)
	ErrRole     = newErrorCode("user role error", 110210)

	// 文件系统错误
	ErrFileNotExists        = newErrorCode("file not exists", 120100)
	ErrFileAlreadyExists    = newErrorCode("file already exists", 120200)
	ErrFileTooLarge         = newErrorCode("file size exceeds limit", 120300)
	ErrUnsupportedMediaType = newErrorCode("file parsing failed", 120400)
	ErrFileIoError          = newErrorCode("file io error", 120500)

	// 检验、签名错误
	ErrInvalidParams = newErrorCode("invalid params", 200100)

	ErrInvalidJSON     = newErrorCode("invalid JSON", 200210)
	ErrInvalidFilename = newErrorCode("invalid filename", 200230)
	ErrInvalidVersion  = newErrorCode("invalid version", 200240)
	WarnCompNotExit    = newErrorCode("comp not exit,please register", 200300)
	WarnCompNotActive  = newErrorCode("comp is approving,please wait", 200310)
	// db model 错误
	ErrAlreadyExists = newErrorCode("object already exists", 300100)
	ErrNotExist      = newErrorCode("object not exist", 300200)
)
