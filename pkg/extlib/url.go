package extlib

import (
	"errors"
	"net/url"
)

const ErrInvalidUrl = "specified url is not valid"

func ValidateURL(u string) error {
	res, err := url.Parse(u)
	if err == nil && res.Scheme == "http" || res.Scheme == "https" {
		return nil
	}

	return errors.New(ErrInvalidUrl)
}
