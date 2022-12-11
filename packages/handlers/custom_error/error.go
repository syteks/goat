package custom_error

import "log"

// Handle handles the display of an error.
// We can pass a custom message to append to the error thrown.
func Handle(err error, customErrorMessage string) {
	if err == nil {
		return
	}

	log.Printf(customErrorMessage+" Error: %s", err)
}
