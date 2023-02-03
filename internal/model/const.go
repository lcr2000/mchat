package model

import "errors"

type Cmd int

const (
	CmdReady     Cmd = 1
	CmdChatEnter Cmd = 2
	CmdChat      Cmd = 3
)

type ErrCode int

const (
	ErrCodeSuccess      ErrCode = 0
	ErrCodeBadParams    ErrCode = 40000
	ErrCodeAccountExist ErrCode = 40001
	ErrCodeServerError  ErrCode = 50000
)

const (
	SYStemName = "SYSTEM"
)

var (
	// ErrBadParams bad params
	ErrBadParams = errors.New("bad params")
	// ErrUnsupported type
	ErrUnsupported = errors.New("unsupported type")
)
