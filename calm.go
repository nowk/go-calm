package calm

// Calm ly recover and return the error
func Calm(fn func()) error {
	var e error
	func() {
		defer func() {
			if err := recover(); err != nil {
				e = err.(error)
			}
		}()

		fn()
	}()

	return e
}
