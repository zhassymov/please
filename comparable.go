package please

import "fmt"

// Empty returns a validation function that checks if the value is empty.
func Empty[T comparable]() Validate[T] {
	return func(value T) error {
		var empty T
		if value == empty {
			return nil
		}
		return fmt.Errorf("%v must be empty", value)
	}
}

// NotEmpty returns a validation function that checks if the value is not empty.
func NotEmpty[T comparable]() Validate[T] {
	return func(value T) error {
		var empty T
		if value != empty {
			return nil
		}
		return fmt.Errorf("%v must not be empty", value)
	}
}

// Equal returns a validation function that checks if the value is equal to the target.
func Equal[T comparable](target T) Validate[T] {
	return func(value T) error {
		if value != target {
			return fmt.Errorf("%v must be equal to %v", value, target)
		}
		return nil
	}
}

// NotEqual returns a validation function that checks if the value is not equal to the target.
func NotEqual[T comparable](target T) Validate[T] {
	return func(value T) error {
		if value == target {
			return fmt.Errorf("%v must not be equal to %v", value, target)
		}
		return nil
	}
}

// OneOf returns a validation function that checks if the value is one of the enum keys.
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

// NotOneOf returns a validation function that checks if the value is not one of the enum keys.
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

// keys returns a slice of keys from the map.
func keys[T comparable](m map[T]bool) []T {
	if len(m) == 0 {
		return nil
	}
	s := make([]T, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// OneIn returns a validation function that checks if the value is one in the enum map.
func OneIn[T comparable](enum map[T]bool) Validate[T] {
	return func(value T) error {
		if _, ok := enum[value]; ok {
			return nil
		}
		return fmt.Errorf("%v must be one in %v", value, keys(enum))
	}
}

// NotOneIn returns a validation function that checks if the value is not one in the enum map.
func NotOneIn[T comparable](enum map[T]bool) Validate[T] {
	return func(value T) error {
		if _, ok := enum[value]; ok {
			return nil
		}
		return fmt.Errorf("%v must not be one in %v", value, keys(enum))
	}
}
