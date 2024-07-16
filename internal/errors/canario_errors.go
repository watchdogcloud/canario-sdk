package errors

type ServerError struct {
	Message string
	Err     error
}

type BadRequestError struct {
	Message string
	Err     error
}

type GatewayError struct {
	Message string
	Err     error
}

// according to Feathers.js v4 (terrier) backend
type CanarioErr struct {
	InternalErrorCode string `json:"internal_error_code"`
	Field             string `json:"field"`
	Description       string `json:"description"`
	Code              string `json:"code"`
}

type CanarioErrorJSON struct {
	ErrorData CanarioErr `json:"error"`
}

func handleError(msg string, err error) string {
	errorMessage := ""
	if msg != "" {
		errorMessage += msg
	}

	if err != nil {
		errorMessage += err.Error()
	}

	return errorMessage
}

// appends error message to message
func (SE *ServerError) Error() string {
	return handleError(SE.Message, SE.Err)
}

func (BR *BadRequestError) Error() string {
	return handleError(BR.Message, BR.Err)
}

func (GE *GatewayError) Error() string {
	return handleError(GE.Message, GE.Err)
}
