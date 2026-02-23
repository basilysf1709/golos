package internal

import "fmt"

type StdoutMode struct{}

func (s *StdoutMode) Deliver(text string) error {
	fmt.Println(text)
	return nil
}
