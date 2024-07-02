package please

import (
	"errors"
	"fmt"
	"slices"
)

// SliceLen returns a validation function that checks whether the length of the slice is equal to the specified number.
func SliceLen[S ~[]E, E any](n int) Validate[S] {
	return func(s S) error {
		if len(s) != n {
			return fmt.Errorf("length must be equal %d", n)
		}
		return nil
	}
}

// SliceMinLen returns a validation function that checks whether the length of the slice is at least the specified number.
func SliceMinLen[S ~[]E, E any](n int) Validate[S] {
	return func(s S) error {
		if len(s) < n {
			return fmt.Errorf("length must be at least %d", n)
		}
		return nil
	}
}

// SliceMaxLen returns a validation function that checks whether the length of the slice is at most the specified number.
func SliceMaxLen[S ~[]E, E any](n int) Validate[S] {
	return func(s S) error {
		if len(s) > n {
			return fmt.Errorf("length must be at most %d", n)
		}
		return nil
	}
}

// SliceLenBetween returns a validation function that checks whether the length of the slice is between the specified numbers.
func SliceLenBetween[S ~[]E, E any](x, y int) Validate[S] {
	return func(s S) error {
		minimal := min(x, y)
		maximal := max(x, y)
		if len(s) < minimal || len(s) > maximal {
			return fmt.Errorf("length must be between %d and %d", minimal, maximal)
		}
		return nil
	}
}

// SliceLenNotBetween returns a validation function that checks whether the length of the slice is not between the specified numbers.
func SliceLenNotBetween[S ~[]E, E any](x, y int) Validate[S] {
	return func(s S) error {
		minimal := min(x, y)
		maximal := max(x, y)
		if len(s) >= minimal && len(s) <= maximal {
			return fmt.Errorf("length must not be between %d and %d", minimal, maximal)
		}
		return nil
	}
}

// SliceContain returns a validation function that checks whether the slice contains the specified value.
func SliceContain[S ~[]E, E comparable](value E) Validate[S] {
	return func(s S) error {
		if !slices.Contains(s, value) {
			return fmt.Errorf("%s must contain %v", s, value)
		}
		return nil
	}
}

// SliceNotContain returns a validation function that checks whether the slice does not contain the specified value.
func SliceNotContain[S ~[]E, E comparable](value E) Validate[S] {
	return func(s S) error {
		if slices.Contains(s, value) {
			return fmt.Errorf("%s must not contain %v", s, value)
		}
		return nil
	}
}

// SliceEach returns a validation function that checks whether each element in the slice satisfies the specified validation functions.
func SliceEach[S ~[]E, E any](opts ...Validate[E]) Validate[S] {
	return func(s S) error {
		errs := make([]error, 0, len(opts))
		for i := range s {
			if err := Join(s[i], opts...); err != nil {
				errs = append(errs, err)
			}
		}
		return errors.Join(errs...)
	}
}
