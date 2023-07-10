package try

import (
	"errors"
	"fmt"
)

// Try0 is a wrapper for a function that returns an error
func Try0(err error) {
	if err != nil {
		panic(err)
	}
}

// Try1 is a wrapper for a function that returns a value and an error
func Try1[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// Try2 is a wrapper for a function that returns two values and an error
func Try2[T1 any, T2 any](value1 T1, value2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return value1, value2
}

// Try3 is a wrapper for a function that returns three values and an error
func Try3[T1 any, T2 any, T3 any](value1 T1, value2 T2, value3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3
}

// Try4 is a wrapper for a function that returns four values and an error
func Try4[T1 any, T2 any, T3 any, T4 any](value1 T1, value2 T2, value3 T3, value4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3, value4
}

// Try5 is a wrapper for a function that returns five values and an error
func Try5[T1 any, T2 any, T3 any, T4 any, T5 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3, value4, value5
}

// Try6 is a wrapper for a function that returns six values and an error
func Try6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, value6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3, value4, value5, value6
}

// Try7 is a wrapper for a function that returns seven values and an error
func Try7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, value6 T6, value7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3, value4, value5, value6, value7
}

// Try8 is a wrapper for a function that returns eight values and an error
func Try8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, value6 T6, value7 T7, value8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3, value4, value5, value6, value7, value8
}

// Try9 is a wrapper for a function that returns nine values and an error
func Try9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, value6 T6, value7 T7, value8 T8, value9 T9, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3, value4, value5, value6, value7, value8, value9
}

// Try10 is a wrapper for a function that returns ten values and an error
func Try10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, value6 T6, value7 T7, value8 T8, value9 T9, value10 T10, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3, value4, value5, value6, value7, value8, value9, value10
}

type tryCatchFinally struct {
	tryCatch
}

type tryCatch struct {
	finally func()
	err     error
}

// Try tries to execute the functions. If any of the functions returns an error,
// the error will be caught and passed to the catch function.
func Try(functions ...func()) (result tryCatchFinally) {
	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(error); ok {
				result.err = v
				return
			}
			result.err = fmt.Errorf("%v", r)
		}
	}()
	for _, function := range functions {
		function()
	}
	return result
}

// CatchMap is a map of errors and their corresponding catch functions.
type CatchMap map[error]func(error)

type anyError struct{}

// AnyError is an error that can be used to catch any error.
var AnyError = anyError{}

func (a anyError) Error() string {
	return "any error"
}

// Catch catches the error and executes the corresponding catch function.
func (t tryCatch) Catch(catchMap CatchMap) {
	defer func() {
		if t.finally != nil {
			t.finally()
		}
	}()
	if t.err != nil {
		for k, v := range catchMap {
			if errors.Is(t.err, k) {
				v(t.err)
				return
			}
		}
		if v, ok := catchMap[AnyError]; ok {
			v(t.err)
			return
		}
	}
}

// Finally executes the final function.
func (t tryCatchFinally) Finally(finally func()) tryCatch {
	t.tryCatch.finally = finally
	return t.tryCatch
}
