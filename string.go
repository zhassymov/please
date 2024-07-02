package please

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// StringLen returns a validation function that checks whether the length of the string is equal to the specified number.
func StringLen(n int) Validate[string] {
	return func(s string) error {
		if len(s) != n {
			return fmt.Errorf("must contain exactly %d characters", n)
		}
		return nil
	}
}

// StringMinLen returns a validation function that checks whether the length of the string is at least the specified number.
func StringMinLen(n int) Validate[string] {
	return func(s string) error {
		if len(s) < n {
			return fmt.Errorf("must contain at least %d characters", n)
		}
		return nil
	}
}

// StringMaxLen returns a validation function that checks whether the length of the string is at most the specified number.
func StringMaxLen(n int) Validate[string] {
	return func(s string) error {
		if len(s) > n {
			return fmt.Errorf("must contain at most %d characters", n)
		}
		return nil
	}
}

// StringLenBetween returns a validation function that checks whether the length of the string is between the specified number.
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

// StringLenNotBetween returns a validation function that checks whether the length of the string is not between the specified number.
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

// StringUTF8 returns a validation function that checks whether the string is a valid UTF-8 string.
func StringUTF8() Validate[string] {
	return func(s string) error {
		if !utf8.ValidString(s) {
			return fmt.Errorf("must be utf-8 valid string")
		}
		return nil
	}
}

// StringRuneCount returns a validation function that checks whether the number of runes in the string is exactly equal to the specified number.
func StringRuneCount(n int) Validate[string] {
	return func(s string) error {
		if utf8.RuneCountInString(s) != n {
			return fmt.Errorf("must contain exactly %d characters", n)
		}
		return nil
	}
}

// StringMinRuneCount returns a validation function that checks whether the number of runes in the string is at least the specified number.
func StringMinRuneCount(n int) Validate[string] {
	return func(s string) error {
		if utf8.RuneCountInString(s) < n {
			return fmt.Errorf("must contain at least %d characters", n)
		}
		return nil
	}
}

// StringMaxRuneCount returns a validation function that checks whether the number of runes in the string is at most the specified number.
func StringMaxRuneCount(n int) Validate[string] {
	return func(s string) error {
		if utf8.RuneCountInString(s) > n {
			return fmt.Errorf("must contain at most %d characters", n)
		}
		return nil
	}
}

// StringRuneCountBetween returns a validation function that checks whether the number of runes in the string is between the specified numbers.
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

// StringRuneCountNotBetween returns a validation function that checks whether the number of runes in the string is not between the specified numbers.
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

// uniqueRuneCount returns the number of unique runes in the string.
func uniqueRuneCount(s string) int {
	unique := make(map[rune]bool)
	for _, char := range s {
		unique[char] = true
	}
	return len(unique)
}

// StringUniqueRuneCount returns a validation function that checks whether the number of unique runes in the string is exactly equal to the specified number.
func StringUniqueRuneCount(n int) Validate[string] {
	return func(s string) error {
		if uniqueRuneCount(s) != n {
			return fmt.Errorf("must contain exactly %d unique characters", n)
		}
		return nil
	}
}

// StringMinUniqueRuneCount returns a validation function that checks whether the number of unique runes in the string is at least the specified number.
func StringMinUniqueRuneCount(n int) Validate[string] {
	return func(s string) error {
		if uniqueRuneCount(s) < n {
			return fmt.Errorf("must contain at least %d unique characters", n)
		}
		return nil
	}
}

// StringMaxUniqueRuneCount returns a validation function that checks whether the number of unique runes in the string is at most the specified number.
func StringMaxUniqueRuneCount(n int) Validate[string] {
	return func(s string) error {
		if uniqueRuneCount(s) < n {
			return fmt.Errorf("must contain at most %d unique characters", n)
		}
		return nil
	}
}

// StringUniqueRuneCountBetween returns a validation function that checks whether the number of unique runes in the string is between the specified numbers.
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

// StringUniqueRuneCountNotBetween returns a validation function that checks whether the number of unique runes in the string is not between the specified numbers.
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

// StringContains returns a validation function that checks whether the string contains the specified substring.
func StringContains(substr string) Validate[string] {
	return func(s string) error {
		if !strings.Contains(s, substr) {
			return fmt.Errorf("must contain %q", substr)
		}
		return nil
	}
}

// StringNotContains returns a validation function that checks whether the string does not contain the specified substring.
func StringNotContains(substr string) Validate[string] {
	return func(s string) error {
		if strings.Contains(s, substr) {
			return fmt.Errorf("must not contain %q", substr)
		}
		return nil
	}
}

// StringHasPrefix returns a validation function that checks whether the string begins with prefix.
func StringHasPrefix(prefix string) Validate[string] {
	return func(s string) error {
		if !strings.HasPrefix(s, prefix) {
			return fmt.Errorf("must contain prefix %q", prefix)
		}
		return nil
	}
}

// StringNotHasPrefix returns a validation function that checks whether the string does not begin with prefix.
func StringNotHasPrefix(prefix string) Validate[string] {
	return func(s string) error {
		if strings.HasPrefix(s, prefix) {
			return fmt.Errorf("must not contain prefix %q", prefix)
		}
		return nil
	}
}

// StringHasSuffix returns a validation function that checks whether the string ends with suffix.
func StringHasSuffix(suffix string) Validate[string] {
	return func(s string) error {
		if !strings.HasSuffix(s, suffix) {
			return fmt.Errorf("must contain suffix %q", suffix)
		}
		return nil
	}
}

// StringNotHasSuffix returns a validation function that checks whether the string does not end with suffix.
func StringNotHasSuffix(suffix string) Validate[string] {
	return func(s string) error {
		if !strings.HasSuffix(s, suffix) {
			return fmt.Errorf("must not contain suffix %q", suffix)
		}
		return nil
	}
}

// StringNumeric returns a validation function that checks whether the string contains only numeric characters.
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

// StringAlpha returns a validation function that checks whether the string contains only alphabet characters.
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

// StringAlphaNumeric returns a validation function that checks whether the string contains only alphanumeric characters.
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

// StringASCII returns a validation function that checks whether the string contains only ASCII printable characters.
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

// StringUnicodeLetters returns a validation function that checks whether the string contains only unicode letters.
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

// StringUnicodeDigits returns a validation function that checks whether the string contains only unicode digits.
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

// StringAllow returns a validation function that checks whether the string contains only allowed characters.
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

// StringNotAllow returns a validation function that checks whether the string does not contain disallowed characters.
func StringNotAllow(charset string) Validate[string] {
	return func(s string) error {
		if strings.ContainsAny(s, charset) {
			return fmt.Errorf("must not contain disallowed characters: %v", charset)
		}
		return nil
	}
}
