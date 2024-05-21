package please

import "fmt"

func Empty[T comparable](value T) error {
	var empty T
	if value == empty {
		return nil
	}
	return fmt.Errorf("%v must be empty", value)
}

func NotEmpty[T comparable](value T) error {
	var empty T
	if value != empty {
		return nil
	}
	return fmt.Errorf("%v must not be empty", value)
}

func Equal[T comparable](target T) Validate[T] {
	return func(value T) error {
		if value != target {
			return fmt.Errorf("%v must be equal to %v", value, target)
		}
		return nil
	}
}

func NotEqual[T comparable](target T) Validate[T] {
	return func(value T) error {
		if value == target {
			return fmt.Errorf("%v must not be equal to %v", value, target)
		}
		return nil
	}
}

func OneOf[T comparable](enum ...T) Validate[T] {
	return func(value T) error {
		for _, e := range enum {
			if value == e {
				return nil
			}
		}
		return fmt.Errorf("%v must be one of %v", value, enum)
	}
}

func NotOneOf[T comparable](enum ...T) Validate[T] {
	return func(value T) error {
		for _, e := range enum {
			if value == e {
				return fmt.Errorf("%v must not be one of %v", value, enum)
			}
		}
		return nil
	}
}

func slice[T comparable](m map[T]bool) []T {
	if len(m) == 0 {
		return nil
	}
	s := make([]T, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func OneIn[T comparable](enum map[T]bool) Validate[T] {
	return func(value T) error {
		if _, ok := enum[value]; ok {
			return nil
		}
		return fmt.Errorf("%v must be one in %v", value, slice(enum))
	}
}

func NotOneIn[T comparable](enum map[T]bool) Validate[T] {
	return func(value T) error {
		if _, ok := enum[value]; ok {
			return nil
		}
		return fmt.Errorf("%v must not be one in %v", value, slice(enum))
	}
}
