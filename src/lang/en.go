package lang

var ValidateField = MessagesMap{
	"username": "username",
}

var Messages = MessagesMap{
	"invalid_data": "Invalid data",
	"not_found":    "Not found",
}

var ValidateErrorsMap = ValidateErrors{
	"required": "The :attr is required",
	"min":      "The :attr must be at least :param characters",
	"max":      "The :attr must be less than :param characters",
}
