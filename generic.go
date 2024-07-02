package please

import (
	"errors"
	"fmt"
)

// Validate is a generic type for validation functions.
type Validate[T any] func(T) error

// WithError returns a new validation function that wraps the original validation function and returns the specified error.
func (v Validate[T]) WithError(cause error) Validate[T] {
	return func(value T) error {
		if err := v(value); err != nil {
			return cause
		}
		return nil
	}
}

// WrapError returns a new validation function that wraps the original validation function and returns the wrapped error.
func (v Validate[T]) WrapError(cause error) Validate[T] {
	return func(value T) error {
		if err := v(value); err != nil {
			return fmt.Errorf("%w: %w", cause, err)
		}
		return nil
	}
}

// Abort returns the first error when executing the validation functions.
func Abort[T any](value T, vs ...Validate[T]) error {
	for _, v := range vs {
		if err := v(value); err != nil {
			return err
		}
	}
	return nil
}

// Collect returns all errors when executing the validation functions.
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

// JoinFunc returns the joined error when executing the validation functions with the specified join function.
func JoinFunc[T any](value T, join func(...error) error, vs ...Validate[T]) error {
	errs := Collect(value, vs...)
	if len(errs) == 0 {
		return nil
	}
	return join(errs...)
}

// Join returns the joined error when executing the validation functions.
func Join[T any](value T, vs ...Validate[T]) error {
	return JoinFunc(value, errors.Join, vs...)
}

// Nothing returns a validation function that always returns nil.
func Nothing[T any]() Validate[T] {
	return func(value T) error {
		return nil
	}
}
