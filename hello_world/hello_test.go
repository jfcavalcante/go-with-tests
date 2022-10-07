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
	//	got := Hello("Chris")
	//	want := "Hello, Chris"
	//
	//	if got != want {
	//		t.Errorf("got %q want %q", got, want)
	//	}

	//subtests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", " ")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Amelie", "French")
		want := "Bonjour, Amelie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Jhon", "Portuguese")
		want := "Olá, Jon"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
