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
