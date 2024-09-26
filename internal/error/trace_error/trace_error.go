package trace_error

import (
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

type TraceError struct {
	Id            string
	Code          string
	UnderlyingErr error
}

func NewTraceError(code string) *TraceError {
	return &TraceError{
		Id:   uuid.NewString(),
		Code: code,
	}
}

func NewTraceErrorWithUnderlying(code string, underlying error) *TraceError {
	return &TraceError{
		Id:            uuid.NewString(),
		Code:          code,
		UnderlyingErr: underlying,
	}
}

func (tr *TraceError) Error() string {
	return fmt.Sprintf("TrErr [%s][%s]", tr.Id, tr.Code)
}

func (tr *TraceError) Generate() *TraceError {
	return NewTraceError(tr.Code)
}

func (tr *TraceError) GenerateAndLog(pattern string, args ...interface{}) *TraceError {
	trErr := NewTraceError(tr.Code)
	slog.Error(pattern, args)
	return trErr
}

func (tr *TraceError) GenerateWithUnderlyingErrAndLog(underlyingErr error, pattern string, args ...interface{}) *TraceError {
	trErr := NewTraceErrorWithUnderlying(tr.Code, underlyingErr)
	return trErr
}

func (tr *TraceError) From(err error) *TraceError {

	lowLevelErr, ok := err.(*TraceError)
	if !ok {
		// This means that there is no "tracing id" from the origin of the error
		slog.Error("TraceError failed to cast the type for error: [%s]", err)
		lowLevelErr = NewTraceError(tr.Code)
	}
	return &TraceError{
		Id:            lowLevelErr.Id,
		Code:          tr.Code,
		UnderlyingErr: err,
	}
}

func (tr *TraceError) Is(err *TraceError) bool {
	if err == nil {
		return false
	}
	return tr.Code == err.Code
}

func (tr *TraceError) Info() string {
	msg := tr.Error()
	if tr.UnderlyingErr != nil {
		msg += fmt.Sprintf("\nInfo: [%s]", tr.UnderlyingErr.Error())
	}
	return msg
}
