package errors

// IErrCode 错误码
type IErrCode interface {
	error
	Code() int
	EnumCode() Enum
}

type errCode struct {
	code Enum
	msg  string
}

// New 创建错误码
func New(retCode Enum, errs ...error) IErrCode {
	errStr := retCode.String()
	// todo如果是debug的话 显示详细错误
	// todo如果是Release的话 直接从map中取
	if errs != nil && len(errs) != 0 {
		errStr = ""
		for _, err := range errs {
			errStr += err.Error()
			errStr += "\n"
		}
	}
	return &errCode{
		code: retCode,
		msg:  errStr,
	}
}

func (e *errCode) Error() string {
	return e.msg
}

func (e *errCode) Code() int {
	return int(e.code)
}

func (e *errCode) EnumCode() Enum {
	return e.code
}
