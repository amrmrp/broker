package errors

type Error string

func (e Error) Error() string{
	panic(e)
}