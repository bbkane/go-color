package color

import "errors"

func Init() error {
	if initialized {
		return errors.New("Init() called more than once")
	}
	setColors()
	initialized = true
	return nil
}
