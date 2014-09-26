package calm

import "fmt"

// Calm ly recover and return the error
func Calm(fn func()) error {
	var e error
	func() {
		defer func() {
			if err := recover(); err != nil {
				switch v := err.(type) {
				case error:
					e = v
				case string:
					e = fmt.Errorf("%s", v)
				}
			}
		}()

		fn()
	}()

	return e
}
