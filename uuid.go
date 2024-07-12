package please

import "github.com/google/uuid"

func UUID() Validate[string] {
	return func(s string) error {
		_, err := uuid.Parse(s)
		if err != nil {
			return err
		}
		return nil
	}
}
