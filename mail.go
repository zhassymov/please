package please

import "net/mail"

// Email returns a validation function for email address.
func Email() Validate[string] {
	return func(s string) error {
		_, err := mail.ParseAddress(s)
		if err != nil {
			return err
		}
		return nil
	}
}
