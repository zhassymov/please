package please

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func StringLen(n int) Validate[string] {
	return func(s string) error {
		if len(s) != n {
			return fmt.Errorf("must contain exactly %d characters", n)
		}
		return nil
	}
}

func StringMinLen(n int) Validate[string] {
	return func(s string) error {
		if len(s) < n {
			return fmt.Errorf("must contain at least %d characters", n)
		}
		return nil
	}
}

func StringMaxLen(n int) Validate[string] {
	return func(s string) error {
		if len(s) > n {
			return fmt.Errorf("must contain at most %d characters", n)
		}
		return nil
	}
}

func StringLenBetween(x, y int) Validate[string] {
	return func(s string) error {
		minimal := min(x, y)
		maximal := max(x, y)
		if len(s) < minimal || len(s) > maximal {
			return fmt.Errorf("must contain from %d to %d characters", minimal, maximal)
		}
		return nil
	}
}

func StringLenNotBetween(x, y int) Validate[string] {
	return func(s string) error {
		minimal := min(x, y)
		maximal := max(x, y)
		if len(s) >= minimal && len(s) <= maximal {
			return fmt.Errorf("must contain up to %d or more than %d characters", minimal, maximal)
		}
		return nil
	}
}

func StringUTF8() Validate[string] {
	return func(s string) error {
		if !utf8.ValidString(s) {
			return fmt.Errorf("must be utf-8 valid string")
		}
		return nil
	}
}

func StringRuneCount(n int) Validate[string] {
	return func(s string) error {
		if utf8.RuneCountInString(s) != n {
			return fmt.Errorf("must contain exactly %d characters", n)
		}
		return nil
	}
}

func StringMinRuneCount(n int) Validate[string] {
	return func(s string) error {
		if utf8.RuneCountInString(s) < n {
			return fmt.Errorf("must contain at least %d characters", n)
		}
		return nil
	}
}

func StringMaxRuneCount(n int) Validate[string] {
	return func(s string) error {
		if utf8.RuneCountInString(s) > n {
			return fmt.Errorf("must contain at most %d characters", n)
		}
		return nil
	}
}

func StringRuneCountBetween(x, y int) Validate[string] {
	return func(s string) error {
		minimal := min(x, y)
		maximal := max(x, y)
		count := utf8.RuneCountInString(s)
		if count < minimal || count > maximal {
			return fmt.Errorf("must contain from %d to %d characters", minimal, maximal)
		}
		return nil
	}
}

func StringRuneCountNotBetween(x, y int) Validate[string] {
	return func(s string) error {
		minimal := min(x, y)
		maximal := max(x, y)
		count := utf8.RuneCountInString(s)
		if count >= minimal && count <= maximal {
			return fmt.Errorf("must contain up to %d or more than %d characters", minimal, maximal)
		}
		return nil
	}
}

func uniqueRuneCount(s string) int {
	unique := make(map[rune]bool)
	for _, char := range s {
		unique[char] = true
	}
	return len(unique)
}

func StringUniqueRuneCount(n int) Validate[string] {
	return func(s string) error {
		if uniqueRuneCount(s) != n {
			return fmt.Errorf("must contain exactly %d unique characters", n)
		}
		return nil
	}
}

func StringMinUniqueRuneCount(n int) Validate[string] {
	return func(s string) error {
		if uniqueRuneCount(s) < n {
			return fmt.Errorf("must contain at least %d unique characters", n)
		}
		return nil
	}
}

func StringMaxUniqueRuneCount(n int) Validate[string] {
	return func(s string) error {
		if uniqueRuneCount(s) < n {
			return fmt.Errorf("must contain at most %d unique characters", n)
		}
		return nil
	}
}

func StringUniqueRuneCountBetween(x, y int) Validate[string] {
	return func(s string) error {
		minimal := min(x, y)
		maximal := max(x, y)
		count := uniqueRuneCount(s)
		if count < minimal || count > maximal {
			return fmt.Errorf("must contain from %d to %d unique characters", minimal, maximal)
		}
		return nil
	}
}

func StringUniqueRuneCountNotBetween(x, y int) Validate[string] {
	return func(s string) error {
		minimal := min(x, y)
		maximal := max(x, y)
		count := uniqueRuneCount(s)
		if count >= minimal && count <= maximal {
			return fmt.Errorf("must contain up to %d or more than %d unique characters", minimal, maximal)
		}
		return nil
	}
}

func StringContains(substr string) Validate[string] {
	return func(s string) error {
		if !strings.Contains(s, substr) {
			return fmt.Errorf("must contain %q", substr)
		}
		return nil
	}
}

func StringNotContains(substr string) Validate[string] {
	return func(s string) error {
		if strings.Contains(s, substr) {
			return fmt.Errorf("must not contain %q", substr)
		}
		return nil
	}
}

func StringHasPrefix(prefix string) Validate[string] {
	return func(s string) error {
		if !strings.HasPrefix(s, prefix) {
			return fmt.Errorf("must contain prefix %q", prefix)
		}
		return nil
	}
}

func StringNotHasPrefix(prefix string) Validate[string] {
	return func(s string) error {
		if strings.HasPrefix(s, prefix) {
			return fmt.Errorf("must not contain prefix %q", prefix)
		}
		return nil
	}
}

func StringHasSuffix(suffix string) Validate[string] {
	return func(s string) error {
		if !strings.HasSuffix(s, suffix) {
			return fmt.Errorf("must contain suffix %q", suffix)
		}
		return nil
	}
}

func StringNotHasSuffix(suffix string) Validate[string] {
	return func(s string) error {
		if !strings.HasSuffix(s, suffix) {
			return fmt.Errorf("must not contain suffix %q", suffix)
		}
		return nil
	}
}

func StringNumeric() Validate[string] {
	return func(s string) error {
		for _, char := range s {
			if char > '0' && char < '9' {
				continue
			}
			return errors.New("must contain only numeric characters")
		}
		return nil
	}
}

func StringAlpha() Validate[string] {
	return func(s string) error {
		for _, char := range s {
			if char > 'A' && char < 'Z' {
				continue
			}
			if char > 'a' && char < 'z' {
				continue
			}
			return errors.New("must contain only alphabet characters")
		}
		return nil
	}
}

func StringAlphaNumeric() Validate[string] {
	return func(s string) error {
		for _, char := range s {
			if char > '0' && char < '9' {
				continue
			}
			if char > 'A' && char < 'Z' {
				continue
			}
			if char > 'a' && char < 'z' {
				continue
			}
			return errors.New("must contain only alphanumeric characters")
		}
		return nil
	}
}

func StringASCII() Validate[string] {
	return func(s string) error {
		for _, char := range s {
			if char >= 33 && char <= 126 {
				continue
			}
			return errors.New("must contain only ascii characters")
		}
		return nil
	}
}

func StringUnicodeLetters() Validate[string] {
	return func(s string) error {
		for _, char := range s {
			if !unicode.IsLetter(char) {
				return errors.New("must contain only unicode letters")
			}
		}
		return nil
	}
}

func StringUnicodeDigits() Validate[string] {
	return func(s string) error {
		for _, char := range s {
			if !unicode.IsDigit(char) {
				return errors.New("must contain only unicode digits")
			}
		}
		return nil
	}
}

func StringAllow(charset string) Validate[string] {
	return func(s string) error {
		for _, char := range s {
			if !strings.ContainsRune(charset, char) {
				return fmt.Errorf("must contain only allowed characters: %q", charset)
			}
		}
		return nil
	}
}

func StringNotAllow(charset string) Validate[string] {
	return func(s string) error {
		if strings.ContainsAny(s, charset) {
			return fmt.Errorf("must not contain disallowed characters: %v", charset)
		}
		return nil
	}
}
