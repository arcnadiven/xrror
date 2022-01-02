package xrror

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Xrror struct {
	Code  string
	Time  string
	Stack string
	Err   string
}

func (x *Xrror) Error() string {
	if x.Code == _DEFAULT_ERROR_CODE {
		return fmt.Sprintf(`%s [%s] %s`, x.Time, x.Stack, x.Err)
	}
	return fmt.Sprintf(`[%s] %s {%s} %s`, x.Code, x.Time, x.Stack, x.Err)
}

func SetPathLayer(n int) {
	_DEFAULT_PATH_LAYER = n
}

func SetStackDepth(n int) {
	_DEFAULT_STACK_DEEPTH = n
}

func SetTimeFormat(str string) {
	_DEFAULT_TIME_FORMAT = str
}

func Error(str string) error {
	return genXrror(_DEFAULT_ERROR_CODE, str, _DEFAULT_STACK_DEEPTH, _DEFAULT_PATH_LAYER)
}

func Errorf(format string, args ...interface{}) error {
	return genXrror(_DEFAULT_ERROR_CODE, fmt.Sprintf(format, args...), _DEFAULT_STACK_DEEPTH, _DEFAULT_PATH_LAYER)
}

func WithStack(err error) error {
	return genXrror(_DEFAULT_ERROR_CODE, err.Error(), _DEFAULT_STACK_DEEPTH, _DEFAULT_STACK_DEEPTH)
}

func ErrorWithCode(code, str string) error {
	return genXrror(code, str, _DEFAULT_STACK_DEEPTH, _DEFAULT_PATH_LAYER)
}

func ErrorfWithCode(code, format string, args ...interface{}) error {
	return genXrror(code, fmt.Sprintf(format, args...), _DEFAULT_STACK_DEEPTH, _DEFAULT_STACK_DEEPTH)
}

func StackWithCode(code string, err error) error {
	return genXrror(code, err.Error(), _DEFAULT_STACK_DEEPTH, _DEFAULT_STACK_DEEPTH)
}

func genXrror(code, str string, depth, pl int) error {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = _DEFAULT_RTC_FILE
		line = _DEFAULT_RTC_LINE
	}
	if pl > 0 {
		pathList := strings.Split(file, string(os.PathSeparator))
		if len(pathList) > pl {
			file = strings.Join(pathList[len(pathList)-pl:], string(os.PathSeparator))
		}
	}
	return &Xrror{
		Code:  code,
		Time:  time.Now().Format(_DEFAULT_TIME_FORMAT),
		Stack: file + `:` + strconv.Itoa(line),
		Err:   str,
	}
}
