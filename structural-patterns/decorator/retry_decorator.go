package decorator

import "time"

type RetryDecorator struct {
	Inner   Http
	Retries int
}

func (d *RetryDecorator) Get(url string) (interface{}, error) {
	var lastErr error
	attempts := d.Retries + 1

	for i := 0; i < attempts; i++ {
		response, err := d.Inner.Get(url)
		if err == nil {
			return response, nil
		}
		lastErr = err
		time.Sleep(time.Millisecond)
	}

	return nil, lastErr
}
