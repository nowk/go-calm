package calm_test

import "errors"
import "testing"
import "github.com/nowk/assert"
import . "github.com/nowk/go-calm"

func TestCalm(t *testing.T) {
	err := Calm(func() {
		panic(errors.New("Return Me!"))
	})
	assert.Equal(t, "Return Me!", err.Error())

	err = Calm(func() {
		// ok!
	})
	assert.Nil(t, err)
}

func TestStringPanic(t *testing.T) {
	err := Calm(func() {
		panic("Return Me!")
	})
	assert.Equal(t, "Return Me!", err.Error())
}
