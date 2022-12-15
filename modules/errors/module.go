package apperror

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func TokenRequiredError() (int, ErrorResponse) {
	return 403, ErrorResponse{
		Code:    403,
		Message: "token required",
	}
}

func FirebaseLibraryError() (int, ErrorResponse) {
	return 500, ErrorResponse{
		Code:    500,
		Message: "firebase setup error",
	}
}

func ForbiddenError() (int, ErrorResponse) {
	return 401, ErrorResponse{
		Code:    401,
		Message: "forbidden request",
	}
}
