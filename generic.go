package please

type Validate[T any] func(T) error

func (v Validate[T]) WithError(cause error) Validate[T] {
	return func(value T) error {
		if err := v(value); err != nil {
			return cause
		}
		return nil
	}
}

func Abort[T any](value T, vs ...Validate[T]) error {
	for _, v := range vs {
		if err := v(value); err != nil {
			return err
		}
	}
	return nil
}
