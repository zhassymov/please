package please

import "errors"

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

func Collect[T any](value T, vs ...Validate[T]) []error {
	if len(vs) == 0 {
		return nil
	}
	errs := make([]error, 0, len(vs))
	for _, v := range vs {
		if err := v(value); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

func JoinFunc[T any](value T, join func(...error) error, vs ...Validate[T]) error {
	errs := Collect(value, vs...)
	if len(errs) == 0 {
		return nil
	}
	return join(errs...)
}

func Join[T any](value T, vs ...Validate[T]) error {
	return JoinFunc(value, errors.Join, vs...)
}

func Nothing[T any]() Validate[T] {
	return func(value T) error {
		return nil
	}
}
