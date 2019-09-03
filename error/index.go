package error

type error struct {
	message string
}

func New(message string) error {
	return error{message}
}
