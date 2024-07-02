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

// Abort returns the first error when executing the validation functions and aborts the execution.
func Abort[T any](value T, opts ...Validate[T]) error {
	for _, v := range opts {
		if err := v(value); err != nil {
			return err
		}
	}
	return nil
}

// Collect collects all errors when executing the validation functions.
func Collect[T any](value T, opts ...Validate[T]) []error {
	if len(opts) == 0 {
		return nil
	}
	errs := make([]error, 0, len(opts))
	for _, v := range opts {
		if err := v(value); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

// JoinFunc joins all errors using the specified join function when executing the validation functions.
func JoinFunc[T any](value T, join func(...error) error, opts ...Validate[T]) error {
	errs := Collect(value, opts...)
	if len(errs) == 0 {
		return nil
	}
	return join(errs...)
}

// Join joins all errors using the errors.Join function when executing the validation functions.
func Join[T any](value T, opts ...Validate[T]) error {
	return JoinFunc(value, errors.Join, opts...)
}

// Nothing returns a validation function that always returns nil.
func Nothing[T any]() Validate[T] {
	return func(value T) error {
		return nil
	}
}
