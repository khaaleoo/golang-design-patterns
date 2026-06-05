package singleton

import "sync"

type singleton struct {
	counter int
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{counter: 0}
	})
	return instance
}

func (s *singleton) Increase() {
	s.counter++
}

func (s *singleton) Get() int {
	return s.counter
}
