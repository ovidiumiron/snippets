package exp

import (
	"log"
	"testing"
)

type Bar interface {
	barMethod()
}

// Want to mock this object. bar implement barMethod so is a Bar
type bar struct {
	baz string
}

func (b *bar) barMethod() {
	log.Println("I don't want to see it in unit tests")
}

type mockbar struct {
	baz string
}

func (f mockbar) barMethod() {
	log.Println("I want to see it in unit tests")
}

type Foo struct {
	data string
	Bar  //interface
}

func (f *Foo) fooMethod() {
	f.barMethod()
}

func TestFooMethod(t *testing.T) {

	//f := Foo{"data", bar{"baz"}}
	f := Foo{"data", mockbar{"baz"}}
	f.fooMethod()
}
