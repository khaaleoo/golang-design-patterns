package decorator

import "fmt"

type LoggingDecorator struct {
	Inner Http
}

func (d *LoggingDecorator) Get(url string) (interface{}, error) {
	fmt.Printf("Requesting %s\n", url)
	response, err := d.Inner.Get(url)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
		return nil, err
	}
	fmt.Printf("Request succeeded: %s\n", url)
	return response, nil
}
