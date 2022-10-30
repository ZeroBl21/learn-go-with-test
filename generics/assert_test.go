package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {

		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "world")
	})

	// AssertEqual(t, 1, "1") // uncomment to see the error
}

func TestStack(t *testing.T) {
	t.Run("integers stack", func(t *testing.T) {
		myStackOfInts := new(StackOfInts)

		// Check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// Add a thing, then check it's notnempty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// Add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)

		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)

		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(StackOfStrings)

		// Check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		// Add a thing, then check it's not empty
		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		// Add another thing, pop it back again
		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")

		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")

		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want interface{}) {
	t.Helper()

	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()

	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()

	if got {
		t.Errorf("got %v, want false", got)
	}
}
