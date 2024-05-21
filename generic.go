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
