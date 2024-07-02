package please

import (
	"cmp"
	"fmt"
)

// Min returns a validation function that checks whether the value is greater or equal than the minimal value.
func Min[T cmp.Ordered](minimal T) Validate[T] {
	return func(value T) error {
		if value < minimal {
			return fmt.Errorf("%v must be greater or equal than %v", value, minimal)
		}
		return nil
	}
}

// Max returns a validation function that checks whether the value is less or equal than the maximal value.
func Max[T cmp.Ordered](maximal T) Validate[T] {
	return func(value T) error {
		if value > maximal {
			return fmt.Errorf("%v must be less or equal than %v", value, maximal)
		}
		return nil
	}
}

// Between returns a validation function that checks whether the value is between the minimal and maximal values.
func Between[T cmp.Ordered](x, y T) Validate[T] {
	return func(value T) error {
		minimal := min(x, y)
		maximal := max(x, y)
		if value < minimal || value > maximal {
			return fmt.Errorf("%v must be between %v and %v", value, minimal, maximal)
		}
		return nil
	}
}

// NotBetween returns a validation function that checks whether the value is not between the minimal and maximal values.
func NotBetween[T cmp.Ordered](x, y T) Validate[T] {
	return func(value T) error {
		minimal := min(x, y)
		maximal := max(x, y)
		if value >= minimal && value <= maximal {
			return fmt.Errorf("%v must not be between %v and %v", value, minimal, maximal)
		}
		return nil
	}
}
