package retry

import (
	"errors"
	"time"
)

func HandleFunc(callback func() error, attempts int, retyableErrors []error) error {
	var err error

loop:
	for i := 0; i < attempts; i++ {
		if i != 0 {
			time.Sleep(time.Second * time.Duration((i-1)*2+1))
		}
		err = callback()

		if err == nil {
			return nil
		}

		if len(retyableErrors) == 0 {
			continue loop
		}

		for _, e := range retyableErrors {
			if errors.Is(err, e) {
				continue loop
			}
		}

		return err
	}

	return err
}
