package exceptions

import "errors"

var (
	//ErrConfigRead 配置读取失败
	ErrConfigRead = errors.New("read config error")

	//ErrParseResult
	ErrParseResult = errors.New("result parse error")
)
