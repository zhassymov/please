package please

import (
	"cmp"
	"fmt"
)

func Min[T cmp.Ordered](minimal T) Validate[T] {
	return func(value T) error {
		if value < minimal {
			return fmt.Errorf("%v must be greater or equal than %v", value, minimal)
		}
		return nil
	}
}

func Max[T cmp.Ordered](maximal T) Validate[T] {
	return func(value T) error {
		if value > maximal {
			return fmt.Errorf("%v must be less or equal than %v", value, maximal)
		}
		return nil
	}
}

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
