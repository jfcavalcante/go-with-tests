// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
package main

import "testing"

//	func TestHello(t *testing.T) {
//		got := Hello()
//		want := "Hllo, world"
//
//		if got != want {
//			// Print out a message and fail the test
//			// For tests %q is very useful as it wraps your values in double quotes
//			t.Errorf("got %q want %q", got, want) // %q safely escape a string and add quotes to it
//		}
//	}
func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
