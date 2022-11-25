package apperror

import (
	"errors"
	"fmt"

	"github.com/Sei-Yukinari/gqlgen-todos/src/util/errorcode"
	"golang.org/x/xerrors"
)

type AppError interface {
	error
	Code() errorcode.ErrorCode
	InfoMessage() string
}

type appError struct {
	err         error
	message     string
	frame       xerrors.Frame
	errCode     errorcode.ErrorCode
	infoMessage string
}

func create(msg string) *appError {
	var e appError
	e.message = msg
	e.frame = xerrors.Caller(2)
	return &e
}

func New(msg string) *appError {
	return create(msg)
}

func Errorf(format string, a ...interface{}) *appError {
	return create(fmt.Sprintf(format, a...))
}

func Wrap(err error, msg ...string) *appError {
	var m string
	if len(msg) != 0 {
		m = msg[0]
	}
	e := create(m)
	e.err = err
	return e
}

func Wrapf(err error, format string, args ...interface{}) *appError {
	e := create(fmt.Sprintf(format, args...))
	e.err = err
	return e
}

func Unwrap(e error) error {
	var appErr *appError
	if errors.As(e, &appErr) {
		return appErr.err
	}

	return e
}

func AsAppError(err error) *appError {
	if err == nil {
		return nil
	}

	var e *appError
	if errors.As(err, &e) {
		return e
	}
	return nil
}

func (e *appError) SetCode(code errorcode.ErrorCode) *appError {
	e.errCode = code
	return e
}

func (err *appError) Info(infoMessage string) *appError {
	err.infoMessage = infoMessage
	return err
}

func (err *appError) Infof(format string, a ...interface{}) *appError {
	err.infoMessage = fmt.Sprintf(format, a...)
	return err
}

func (e *appError) Error() string {
	if e.err == nil {
		return e.message
	}
	if e.message != "" {
		return e.message + ": " + e.err.Error()
	}
	return e.err.Error()
}

func (e *appError) Unwrap() error {
	return e.err
}

func (e *appError) Code() errorcode.ErrorCode {
	var next *appError = e
	for next.errCode == "" {
		if err := AsAppError(next.err); err != nil {
			next = err
		} else {
			return errorcode.Unknown
		}
	}
	return next.errCode
}

func (e *appError) InfoMessage() string {
	var next *appError = e
	for next.infoMessage == "" {
		if err := AsAppError(next.err); err != nil {
			next = err
		} else {
			return ""
		}
	}
	return next.infoMessage
}

func (e *appError) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
}

func (e *appError) FormatError(p xerrors.Printer) error {
	p.Print(e.Error())
	if p.Detail() {
		e.frame.Format(p)
	}
	return e.err
}
