package errors

import (
	"github.com/dup2X/gopkg/errcode"
	"github.com/dup2X/gopkg/idl"
)

var _ idl.APIErr = new(CrErr)

// CrErr ...
type CrErr struct {
	code errcode.RespCode
	err  error
}

// Code ...
func (pe *CrErr) Code() int {
	return int(pe.code)
}

// Error impl error
func (pe *CrErr) Error() string {
	if pe.err == nil {
		return "ok"
	}
	return pe.err.Error()
}

// GenError ...
func GenError(code errcode.RespCode, err error) *CrErr {
	return &CrErr{
		code: code,
		err:  err,
	}
}
