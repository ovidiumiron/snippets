package exp

import (
	"log"
	"testing"
)

type Bar struct {
	baz string
}

// Want to mock this call
func (b *Bar) barMethod() {
	log.Println("I don't want to see it in unit tests")
}

type Foo struct {
	data string
	Bar
}

func (f *Foo) fooMethod() {
	f.barMethod()
}

func TestFooMethod(t *testing.T) {
	f := Foo{"data", Bar{"baz"}}
	f.fooMethod()
}
