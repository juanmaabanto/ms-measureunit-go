package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown    = ErrorType{"unknown"}
	ErrorTypeBadRequest = ErrorType{"bad-request"}
	ErrorTypeNotFound   = ErrorType{"not-found"}
	ErrorTypeValidation = ErrorType{"Validation Failure"}
)

type ApplicationError struct {
	error     string
	title     string
	errorType ErrorType
}

func (app ApplicationError) Error() string {
	return app.error
}

func (app ApplicationError) Title() string {
	return app.title
}

func (app ApplicationError) ErrorType() ErrorType {
	return app.errorType
}

func NewApplicationError(error string, title string) ApplicationError {
	return ApplicationError{
		error:     error,
		title:     title,
		errorType: ErrorTypeUnknown,
	}
}

func NewBadRequestError(error string) ApplicationError {
	return ApplicationError{
		error:     error,
		title:     "Bad Request",
		errorType: ErrorTypeBadRequest,
	}
}

func NewNotFoundError(error string) ApplicationError {
	return ApplicationError{
		error:     error,
		title:     "Not Found",
		errorType: ErrorTypeNotFound,
	}
}
