package exp

import (
	"log"
	"testing"
)

type IBar interface {
	barMethod() int
}

type Bar struct {
	baz string
}

// Want to mock this call
func (b *Bar) barMethod() int {
	log.Println("I don't want to see it in unit tests")
	return 10
}

type mockbar struct{}

func (f mockbar) barMethod() int {
	log.Println("I want to see it in unit tests")
	return 10
}

type Foo struct {
	data string
	IBar
}

func (f *Foo) fooMethod() {
	log.Println(f.barMethod() + 1)
}

func TestFooMethod(t *testing.T) {
	//f := Foo{"data", Bar{"baz"}}
	f := Foo{"data", mockbar{}}
	f.fooMethod()
}
