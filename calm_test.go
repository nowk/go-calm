package calm

import "errors"
import "testing"
import "gopkg.in/nowk/assert.v2"

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
