package main

import "log"

// Must kills the app iff last argument is a non-nil error.
func Must(args ...interface{}) {
	err, ok := args[len(args)-1].(error)
	if ok && err != nil {
		log.Fatal(err)
	}
}
