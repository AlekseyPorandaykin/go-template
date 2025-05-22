package terror

const (
	ErrValidation      = "validation_error"
	ErrInternal        = "internal_error"
	ErrStorage         = "storage_error"
	ErrExternalService = "external_service_error"
	ErrLogic           = "logic_error"
)

type Error interface {
	Error() string
	Code() string
	Cause() error
}

type err struct {
	code    string
	message string
	cause   error
	params  map[string]string
}

func New(code, message string) Error {
	return err{
		code:    code,
		message: message,
	}
}

func NewErrValidation(message string) Error {
	return New(ErrValidation, message)
}

func NewErrInternal(message string) Error {
	return New(ErrInternal, message)
}
func NewErrStorage(message string) Error {
	return New(ErrStorage, message)
}
func NewErrLogic(message string) Error {
	return New(ErrLogic, message)
}
func NewErrExternalService(message string) Error {
	return New(ErrExternalService, message)
}

func WrapErrValidation(cause error, message string) Error {
	return Wrap(cause, ErrValidation, message)
}

func WrapErrInternal(cause error, message string) Error {
	return Wrap(cause, ErrInternal, message)
}

func WrapErrStorage(cause error, message string) Error {
	return Wrap(cause, ErrStorage, message)
}

func WrapErrLogic(cause error, message string) Error {
	return Wrap(cause, ErrLogic, message)
}
func WrapErrExternalService(cause error, message string) Error {
	return Wrap(cause, ErrExternalService, message)
}

func Wrap(cause error, code, message string) Error {
	return err{
		code:    code,
		message: message,
		cause:   cause,
	}
}

func With() {

}

func (e err) Error() string {
	res := e.message
	if e.cause != nil {
		res += ": " + e.cause.Error()

	}
	return res
}

func (e err) Code() string {
	return e.code
}

func (e err) Cause() error {
	return e.cause
}
