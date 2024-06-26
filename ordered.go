package please

import (
	"cmp"
	"fmt"
)

func Min[T cmp.Ordered](min T) Validate[T] {
	return func(value T) error {
		if value < min {
			return fmt.Errorf("%v must be greater or equal than %v", value, min)
		}
		return nil
	}
}

func Max[T cmp.Ordered](max T) Validate[T] {
	return func(value T) error {
		if value > max {
			return fmt.Errorf("%v must be less or equal than %v", value, max)
		}
		return nil
	}
}

func Between[T cmp.Ordered](min, max T) Validate[T] {
	return func(value T) error {
		if value < min || value > max {
			return fmt.Errorf("%v must be between %v and %v", value, min, max)
		}
		return nil
	}
}

func NotBetween[T cmp.Ordered](min, max T) Validate[T] {
	return func(value T) error {
		if value >= min && value <= max {
			return fmt.Errorf("%v must not be between %v and %v", value, min, max)
		}
		return nil
	}
}
