package trace_error

var (
	ErrUnexpected = NewTraceError("ERR_UNEXPECTED")
	ErrNotFound   = NewTraceError("ERR_NOT_FOUND")
)
