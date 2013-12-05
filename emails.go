package sesame

import (
	"errors"
	"strings"
)

// ValidateEmail validates a given email address, possibly returning an error.
// It does a very simple check: for the presence of "@". Anything further would
// be silly, since proper email address parsing is hideously complicated. The
// easiest way to validate an email address is to try and send an email to it,
// and that kind of account activation is an outstanding TODO for this repo.
func ValidateEmail(address string) error {
	if !strings.Contains(address, "@") {
		return errors.New("email address is not valid: no @")
	}

	return nil
}
