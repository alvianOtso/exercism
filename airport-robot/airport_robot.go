package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(visitor string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", greeter.LanguageName(), greeter.Greet(visitor))
}

type Italian struct {
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(visitor string) string {
	return "Ciao " + visitor
}


type Portuguese struct {
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(visitor string) string {
	return "Olá " + visitor
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
