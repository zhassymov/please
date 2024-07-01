package please

import (
	"errors"
	"fmt"
	"slices"
)

func SliceLen[S ~[]E, E any](n int) Validate[S] {
	return func(s S) error {
		if len(s) != n {
			return fmt.Errorf("length must be equal %d", n)
		}
		return nil
	}
}

func SliceMinLen[S ~[]E, E any](n int) Validate[S] {
	return func(s S) error {
		if len(s) < n {
			return fmt.Errorf("length must be at least %d", n)
		}
		return nil
	}
}

func SliceMaxLen[S ~[]E, E any](n int) Validate[S] {
	return func(s S) error {
		if len(s) > n {
			return fmt.Errorf("length must be at most %d", n)
		}
		return nil
	}
}

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

func SliceContain[S ~[]E, E comparable](value E) Validate[S] {
	return func(s S) error {
		if !slices.Contains(s, value) {
			return fmt.Errorf("%s must contain %v", s, value)
		}
		return nil
	}
}

func SliceNotContain[S ~[]E, E comparable](value E) Validate[S] {
	return func(s S) error {
		if slices.Contains(s, value) {
			return fmt.Errorf("%s must not contain %v", s, value)
		}
		return nil
	}
}

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
